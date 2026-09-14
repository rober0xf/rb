package password

import "github.com/charmbracelet/lipgloss"

var bocaBlue = lipgloss.Color("#0000cd")
var bocaGold = lipgloss.Color("#fdff00")

var boxStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(bocaBlue).
	Padding(1, 2).
	Width(50)

var (
	appNameStyle = lipgloss.NewStyle().
			Bold(true).
			Background(bocaBlue).
			Foreground(bocaGold).
			Padding(0, 2).
			Align(lipgloss.Center).
			Width(46)

	labelStyle = lipgloss.NewStyle().
			Bold(true).
			Background(bocaBlue).
			Foreground(bocaGold).
			Padding(0, 1)

	faintStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8fa0ff")).
			Faint(true)

	itemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252")).
			PaddingLeft(1)

	selectedItemStyle = lipgloss.NewStyle().
				Bold(true).
				Background(bocaBlue).
				Foreground(bocaGold).
				Padding(0, 1)

	indicatorStyle = lipgloss.NewStyle().
			Bold(true).
			Background(bocaBlue).
			Foreground(bocaGold).
			Padding(0, 1).
			MarginRight(1)
)
