package password

import "github.com/charmbracelet/lipgloss"

var blue = lipgloss.Color("#0000cd")
var yellow = lipgloss.Color("#fdff00")
var lightblue = lipgloss.Color("#61afef")

var boxStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(blue).
	Padding(1, 2).
	Width(50)

var (
	appNameStyle = lipgloss.NewStyle().
			Bold(true).
			Background(blue).
			Foreground(yellow).
			Padding(0, 2).
			Align(lipgloss.Center).
			Width(46)

	labelStyle = lipgloss.NewStyle().
			Bold(true).
			Background(blue).
			Foreground(yellow).
			Padding(0, 1)

	faintStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8fa0ff")).
			Faint(true)

	itemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252")).
			PaddingLeft(1)

	selectedItemStyle = lipgloss.NewStyle().
				Bold(true).
				Background(blue).
				Foreground(yellow).
				Padding(0, 1)

	indicatorStyle = lipgloss.NewStyle().
			Bold(true).
			Background(blue).
			Foreground(yellow).
			Padding(0, 1).
			MarginRight(1)

	confirmDeleteStyle = lipgloss.NewStyle().
				Foreground(yellow).
				Bold(true).
				Padding(0, 1).
				Width(46).
				Align(lipgloss.Center)

	dirStyle = lipgloss.NewStyle().
			Foreground(lightblue)

	tickStyle = lipgloss.NewStyle().
			Foreground(yellow)

	hintStyle = lipgloss.NewStyle().
			Bold(true).
			Background(blue).
			Foreground(yellow).
			Padding(0, 1).
			Width(46)
)
