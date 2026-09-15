package password

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	state            viewState
	bodyField        bodyField
	current          Password
	textinput        textinput.Model
	descriptionInput textinput.Model
	passwordInput    textinput.Model
	passwordsCommand *PasswordManager
	confirmDelete    string
	currentDir       string
	cursor           int
	entries          []os.DirEntry
}

func NewModel(manager *PasswordManager) (Model, error) {
	m := Model{
		state:            listView,
		currentDir:       manager.storeDir,
		cursor:           0,
		passwordsCommand: manager,
		textinput:        textinput.New(),
		descriptionInput: textinput.New(),
		passwordInput:    textinput.New(),
	}

	if err := m.loadDirectory(); err != nil {
		return Model{}, fmt.Errorf("failed to load password store: %w", err)
	}

	return m, nil
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
			case keyLeft:
				if m.currentDir == "" || m.currentDir == m.passwordsCommand.storeDir {
					return m, nil
				}

				m.currentDir = filepath.Dir(m.currentDir)
				m.cursor = 0
				_ = m.loadDirectory()
				return m, nil

			case keyRight:
				if len(m.entries) == 0 {
					return m, nil
				}

				entry := m.entries[m.cursor]
				if entry.IsDir() {
					m.currentDir = filepath.Join(m.currentDir, entry.Name())
					m.cursor = 0

					if err := m.loadDirectory(); err != nil {
						return m, nil
					}
				}
				return m, nil

			case "d":
				if len(m.entries) == 0 {
					break
				}

				entry := m.entries[m.cursor]
				if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".gpg") {
					break
				}

				m.confirmDelete = m.entryTitle(entry)
				m.state = confirmDeleteView
				return m, nil

			case "q", "ctrl+c":
				return m, tea.Quit

			case keyUp, "k":
				if m.cursor > 0 {
					m.cursor--
				}

			case keyDown, "j":
				if m.cursor < len(m.entries)-1 {
					m.cursor++
				}

			case "n":
				m.textinput.SetValue("")
				m.textinput.Focus()
				m.current = Password{}
				m.state = TitleView

				return m, nil

			case "c":
				if len(m.entries) == 0 {
					break
				}

				entry := m.entries[m.cursor]
				if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".gpg") {
					break
				}

				title := m.entryTitle(entry)

				return m, func() tea.Msg {
					_ = m.passwordsCommand.CopyPassword(title)
					return nil
				}

			case "enter":
				if len(m.entries) == 0 {
					break
				}

				entry := m.entries[m.cursor]
				if entry.IsDir() {
					m.currentDir = filepath.Join(m.currentDir, entry.Name())
					m.cursor = 0

					if err := m.loadDirectory(); err != nil {
						return m, nil
					}
					return m, nil
				}

				if !strings.HasSuffix(entry.Name(), ".gpg") {
					break
				}

				m.current.Title = m.entryTitle(entry)

				password, err := m.passwordsCommand.ShowPassword(m.current.Title)
				if err != nil {
					return m, tea.Quit
				}

				m.current = password
				m.descriptionInput.SetValue(m.current.Description)
				m.passwordInput.SetValue(m.current.Password)
				m.passwordInput.EchoMode = textinput.EchoPassword
				m.bodyField = descriptionField
				m.descriptionInput.Focus()
				m.state = BodyView
			}

		case TitleView:
			switch key {
			case "enter":
				title := m.textinput.Value()
				if title != "" {
					m.current.Title = m.passTitle(title)
					m.descriptionInput.SetValue("")
					m.passwordInput.SetValue("")
					m.passwordInput.EchoMode = textinput.EchoPassword
					m.textinput.Blur()
					m.bodyField = descriptionField
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
					return m, nil
				}

				m.current = Password{}

				m.descriptionInput.Blur()
				m.passwordInput.Blur()

				if err := m.loadDirectory(); err != nil {
					return m, nil
				}

				if m.cursor >= len(m.entries) {
					m.cursor = max(0, len(m.entries)-1)
				}

				m.state = listView

				return m, nil

			case "esc":
				m.descriptionInput.Blur()
				m.passwordInput.Blur()
				m.state = listView

				return m, nil
			}

		case confirmDeleteView:
			switch key {
			case "y":
				return m.DeletePassword(m.confirmDelete)

			case "n", "esc":
				m.confirmDelete = ""
				m.state = listView
				return m, nil

			case "q", "ctrl+c":
				return m, tea.Quit
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

func (m Model) DeletePassword(title string) (tea.Model, tea.Cmd) {
	_ = m.passwordsCommand.DeletePassword(title)

	if err := m.loadDirectory(); err != nil {
		m.currentDir = m.passwordsCommand.storeDir
		m.cursor = 0
		if err := m.loadDirectory(); err != nil {
			return m, tea.Quit
		}
	}

	m.confirmDelete = ""
	m.state = listView

	if m.cursor >= len(m.entries) {
		m.cursor = max(0, len(m.entries)-1)
	}

	return m, nil
}

func (m *Model) loadDirectory() error {
	entries, err := os.ReadDir(m.currentDir)
	if err != nil {
		return err
	}

	m.entries = entries[:0]

	for _, entry := range entries {
		if entry.Name() == ".gpg-id" {
			continue
		}

		m.entries = append(m.entries, entry)
	}

	return nil
}
