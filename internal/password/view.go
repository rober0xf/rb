package password

import (
	"strings"
)

func (m Model) View() string {
	var b strings.Builder

	b.WriteString(appNameStyle.Render("PASSWORD MANAGER"))
	b.WriteString("\n\n")

	switch m.state {
	case TitleView:
		b.WriteString("Title: \n")
		b.WriteString(m.textinput.View())
		b.WriteString("\n\n")
		b.WriteString(faintStyle.Render("enter - save, esc - discard"))

	case BodyView:
		b.WriteString("Description: \n")
		b.WriteString(m.descriptionInput.View())
		b.WriteString("\n\n")
		b.WriteString("Password: \n")
		b.WriteString(m.passwordInput.View())
		b.WriteString("\n\n")
		b.WriteString(faintStyle.Render("ctrl+s - save, esc - discard, space - toggle reveal"))

	case listView:
		for i, password := range m.passwords {
			prefix := " "
			if i == m.listIndex {
				prefix = "> "
			}
			b.WriteString("\n")

			b.WriteString(enumeratoStyle.Render(prefix))
			b.WriteString(password.Title)
		}

		b.WriteString("\n")
		b.WriteString(faintStyle.Render("n - new password, q - quit"))
	}

	return b.String()
}
