package password

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type viewState int
type bodyField int

const (
	listView viewState = iota
	TitleView
	BodyView
)

const (
	descriptionField bodyField = iota
	passwordField
)

type Model struct {
	state            viewState
	bodyField        bodyField
	passwords        []Password
	current          Password
	listIndex        int
	textinput        textinput.Model
	descriptionInput textinput.Model
	passwordInput    textinput.Model
	passwordsCommand *PasswordManager
}

func NewModel(manager *PasswordManager) (Model, error) {
	passwords, err := manager.ListPasswords()
	if err != nil {
		return Model{}, fmt.Errorf("failed to show passwords: %w", err)
	}

	return Model{
		state:            listView,
		passwords:        passwords,
		passwordsCommand: manager,
		textinput:        textinput.New(),
		descriptionInput: textinput.New(),
		passwordInput:    textinput.New(),
	}, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) SavePassword() error {
	if err := m.passwordsCommand.SavePassword(m.current); err != nil {
		return fmt.Errorf("failed to save password: %w", err)
	}
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String() // up, down, ctrl+c

		switch m.state {
		case listView:
			switch key {
			case "q", "ctrl+c":
				return m, tea.Quit

			case "up", "k":
				if m.listIndex > 0 {
					m.listIndex--
				}

			case "down", "j":
				if m.listIndex < len(m.passwords)-1 {
					m.listIndex++
				}

			case "n":
				m.textinput.SetValue("")
				m.textinput.Focus()
				m.current = Password{}
				m.state = TitleView

				return m, nil

			case "enter":
				if len(m.passwords) == 0 {
					break
				}
				m.current = m.passwords[m.listIndex]

				password, err := m.passwordsCommand.ShowPassword(m.current.Title)
				if err != nil {
					return m, tea.Quit
				}

				m.current = password
				m.descriptionInput.SetValue(m.current.Description)
				m.passwordInput.SetValue(m.current.Password)
				m.passwordInput.EchoMode = textinput.EchoPassword
				m.descriptionInput.Focus()
				m.state = BodyView
			}

		case TitleView:
			switch key {
			case "enter":
				title := m.textinput.Value()
				if title != "" {
					m.current.Title = title
					m.descriptionInput.SetValue("")
					m.passwordInput.SetValue("")
					m.passwordInput.EchoMode = textinput.EchoPassword
					m.descriptionInput.Focus()
					m.state = BodyView
				}

			case "esc":
				m.textinput.Blur()
				m.state = listView
			}

		case BodyView:
			switch key {
			case "tab":
				// change section focus
				if m.bodyField == descriptionField {
					m.descriptionInput.Blur()
					m.passwordInput.Focus()
					m.bodyField = passwordField
				} else {
					m.passwordInput.Blur()
					m.descriptionInput.Focus()
					m.bodyField = descriptionField
				}

				return m, nil

			case " ":
				if m.bodyField == passwordField {
					if m.passwordInput.EchoMode == textinput.EchoPassword {
						m.passwordInput.EchoMode = textinput.EchoNormal
					} else {
						m.passwordInput.EchoMode = textinput.EchoPassword
					}

					return m, nil
				}

			case "ctrl+s":
				m.current.Description = m.descriptionInput.Value()
				m.current.Password = m.passwordInput.Value()

				if err := m.passwordsCommand.SavePassword(m.current); err != nil {
					return m, tea.Quit
				}

				passwords, err := m.passwordsCommand.ListPasswords()
				if err != nil {
					return m, tea.Quit
				}

				m.passwords = passwords
				m.current = Password{}

				m.descriptionInput.Blur()
				m.passwordInput.Blur()

				m.state = listView

				return m, nil

			case "esc":
				m.descriptionInput.Blur()
				m.passwordInput.Blur()
				m.state = listView

				return m, nil
			}
		}
	}

	switch m.state {
	case TitleView:
		m.textinput, cmd = m.textinput.Update(msg)

	case BodyView:
		if m.bodyField == descriptionField {
			m.descriptionInput, cmd = m.descriptionInput.Update(msg)
		} else {
			m.passwordInput, cmd = m.passwordInput.Update(msg)
		}
	}

	return m, cmd
}
