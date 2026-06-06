package app

import (
	tea "charm.land/bubbletea/v2"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	m.loadActiveBuffer()
	defer func() { (&m).saveActiveBuffer() }()

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.viewWidth = msg.Width
		m.viewHeight = msg.Height

	case tea.MouseClickMsg:
		if msg.Button == tea.MouseLeft {
			m.selecting = false
			m.lastClickX = msg.X
			m.lastClickY = msg.Y
			m.setCursorFromClick(msg.X, msg.Y-1)
		}

	case tea.PasteMsg:
		m = m.handlePasteText(msg.String())

	case tea.KeyPressMsg:
		switch msg.String() {

		// new tab
		case "ctrl+n":
			m.saveActiveBuffer()
			m.buffers = append(m.buffers, buffer{lines: []string{""}, saved: true})
			m.activeBuffer = len(m.buffers) - 1
			m.loadActiveBuffer()

		// next tab
		case "alt+l", "f7":
			if len(m.buffers) > 0 {
				m.saveActiveBuffer()
				m.activeBuffer = (m.activeBuffer + 1) % len(m.buffers)
				m.loadActiveBuffer()
			}

		// prev tab
		case "alt+h", "f6":
			if len(m.buffers) > 0 {
				m.saveActiveBuffer()
				m.activeBuffer = (m.activeBuffer - 1 + len(m.buffers)) % len(m.buffers)
				m.loadActiveBuffer()
			}


		// Quit
		case "esc":
			return m, tea.Quit

		// Copy
		case "ctrl+c":
			m = m.handleCopy()

		// Paste
		case "alt+v":
			m = m.handlePaste()

		// Save
		case "ctrl+s":
			m = m.handleSave()

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
