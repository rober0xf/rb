package password

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	var b strings.Builder

	header := appNameStyle.Width(46).Render("PASSWORD MANAGER")

	b.WriteString(header)
	b.WriteString("\n\n")

	switch m.state {
	case TitleView:
		b.WriteString(labelStyle.Render("Title:"))
		b.WriteString("\n")
		b.WriteString(m.textinput.View())
		b.WriteString("\n\n")
		b.WriteString(indicatorStyle.Render("enter - save, esc - discard"))

	case BodyView:
		b.WriteString(labelStyle.Render("Description:"))
		b.WriteString("\n")
		b.WriteString(m.descriptionInput.View())
		b.WriteString("\n\n")
		b.WriteString(labelStyle.Render("Password:"))
		b.WriteString("\n")
		b.WriteString(m.passwordInput.View())
		b.WriteString("\n\n")
		b.WriteString(indicatorStyle.Render("ctrl+s - save, esc - discard, space - toggle reveal"))

	case listView:
		for i, password := range m.passwords {
			b.WriteString("\n")

			if i == m.listIndex {
				b.WriteString(selectedItemStyle.Render("> " + password.Title))
			} else {
				b.WriteString(itemStyle.Render("  " + password.Title))
			}
		}

		b.WriteString("\n\n")
		b.WriteString(indicatorStyle.Render("n - new password, q - quit, d - delete"))

	case confirmDeleteView:
		fmt.Fprintf(&b, "Delete %q?\n\n", m.confirmDelete)
		b.WriteString(indicatorStyle.Render("y - confirm, n - cancel, q - quit"))
	}

	return boxStyle.Render(b.String())
}
