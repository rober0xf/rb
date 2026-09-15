package password

import (
	"strings"
)

// show stacked
func writeHints(b *strings.Builder, hints ...string) {
	for i, h := range hints {
		b.WriteString(hintStyle.Render(h))

		if i < len(hints)-1 {
			b.WriteString("\n")
		}
	}
}

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
		writeHints(&b, "enter - save", "esc - discard")

	case BodyView:
		b.WriteString(labelStyle.Render("Description:"))
		b.WriteString("\n")
		b.WriteString(m.descriptionInput.View())
		b.WriteString("\n\n")
		b.WriteString(labelStyle.Render("Password:"))
		b.WriteString("\n")
		b.WriteString(m.passwordInput.View())
		b.WriteString("\n\n")
		writeHints(&b, "ctrl+s - save", "esc - discard", "space - toggle reveal")

	case listView:
		if dir := m.displayDir(); dir != "" {
			b.WriteString(dirStyle.Render(dir))
			b.WriteString("\n\n")
		}

		for i, entry := range m.entries {
			prefix := "  "

			if i == m.cursor {
				prefix = "> "
				b.WriteString(tickStyle.Render(prefix))
			}

			name := entry.Name()
			if entry.IsDir() {
				name = dirStyle.Render(name + "/")
			} else {
				name = strings.TrimSuffix(name, ".gpg")
				name = itemStyle.Render(name)
			}

			b.WriteString(name)
			b.WriteString("\n")
		}

		b.WriteString("\n")
		writeHints(&b, "n - new password", "q - quit", "d - delete", "c - copy to clipboard")

	case confirmDeleteView:
		b.WriteString(confirmDeleteStyle.Render(m.confirmDelete))
		b.WriteString("\n\n")
		writeHints(&b, "y - confirm", "n - cancel", "q - quit")
	}

	return boxStyle.Render(b.String())
}
