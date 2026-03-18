package app

import (
	tea "charm.land/bubbletea/v2"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.viewWidth = msg.Width
		m.viewHeight = msg.Height

	case tea.MouseClickMsg:
		if msg.Button == tea.MouseLeft {
			m.lastClickX = msg.X
			m.lastClickY = msg.Y
			m.setCursorFromClick(msg.X, msg.Y-1)
		}

	case tea.KeyPressMsg:
		switch msg.String() {

		// Quit
		case "ctrl+c":
			return m, tea.Quit

		// Movement
		case "up", "down", "right", "left", "home", "end":
			m = m.handleMovement(msg.String())

		case "shift+left", "shift+right", "shift+up", "shift+down":
			m = m.handleSelection(msg.String())

		// Regular character typing
		default:
			m = m.handleEditing(msg)

		}
	}

	(&m).scrollToCursor()
	return m, nil
}

func (m *model) scrollToCursor() {
	if m.cursorRow < m.offsetRow {
		m.offsetRow = m.cursorRow
	}

	if m.cursorRow >= m.offsetRow+m.viewHeight-2 {
		m.offsetRow = m.cursorRow - m.viewHeight + 3
	}
}

func (m *model) setCursorFromClick(screenCol, screenRow int) {

	screenRowSoFar := 0

	for lineIdx, line := range m.lines[m.offsetRow:] {
		actualLineIdx := m.offsetRow + lineIdx

		rowsThisLineTakes := 1
		if len(line) > 0 {
			rowsThisLineTakes = (len(line) + m.viewWidth - 1) / m.viewWidth
		}

		// does the clicked row fall inside this line?
		if screenRow < screenRowSoFar+rowsThisLineTakes {
			m.cursorRow = actualLineIdx // found the line

			chunkIndex := screenRow - screenRowSoFar

			chunkStart := chunkIndex * m.viewWidth

			realCol := chunkStart + screenCol

			if realCol > len(line) {
				realCol = len(line)
			}

			m.cursorCol = realCol
			return
		}

		screenRowSoFar += rowsThisLineTakes
	}
}
