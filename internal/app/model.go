package app

import tea "charm.land/bubbletea/v2"

type model struct {
	// Content
	content    string
	lines      []string

	// Cursor Position
	cursorRow  int
	cursorCol  int

	// Viewport (for scrolling)
	offsetRow  int
	offsetCol  int
	viewWidth  int
	viewHeight int

	// File metadata
	filename string
	saved    bool


}

func InitialModel() model {
	return model{
		lines: []string{""},
		cursorRow: 0,
		cursorCol: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}