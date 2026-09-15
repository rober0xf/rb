package password

type viewState int
type bodyField int

const (
	listView viewState = iota
	TitleView
	BodyView
	confirmDeleteView
)

const (
	descriptionField bodyField = iota
	passwordField
)

const (
	keyUp    = "up"
	keyDown  = "down"
	keyLeft  = "left"
	keyRight = "right"
)
