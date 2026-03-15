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

		case "home":
			m.cursorCol = 0

		case "end":
			m.cursorCol = len(m.lines[m.cursorRow])

		// Enter (insert a new line)
		case "enter":
			currentLine := m.lines[m.cursorRow]
			before := currentLine[:m.cursorCol] // everything left of cursor
			after := currentLine[m.cursorCol:]  // everything right of cursor

			// replace the current line with the left part
			// then insert the right part as a new line below
			m.lines[m.cursorRow] = before
			newLines := make([]string, len(m.lines)+1)
			copy(newLines, m.lines[:m.cursorRow+1])
			newLines[m.cursorRow+1] = after
			copy(newLines[m.cursorRow+2:], m.lines[m.cursorRow+1:])
			m.lines = newLines

			// move cursor to the start of the new lin
			m.cursorRow++
			m.cursorCol = 0

		// Backspace (delete character before cursor)
		case "backspace":
			if m.cursorCol > 0 {
				// there's a character to the left that needs to be deleted
				line := m.lines[m.cursorRow]
				m.lines[m.cursorRow] = line[:m.cursorCol-1] + line[m.cursorCol:]
				m.cursorCol--
			} else if m.cursorRow > 0 {
				// Cusrsor is at the start of the line - merge this line with the one above
				above := m.lines[m.cursorRow-1]
				current := m.lines[m.cursorRow]
				m.cursorCol = len(above)
				m.lines[m.cursorRow-1] = above + current
				m.lines = append(m.lines[:m.cursorRow], m.lines[m.cursorRow+1:]...)
				m.cursorRow--
			}

		case "space":
			line := m.lines[m.cursorRow]
			m.lines[m.cursorRow] = line[:m.cursorCol] + " " + line[m.cursorCol:]
			m.cursorCol++

			// Arrow keys
		case "left":
			if m.cursorCol > 0 {
				m.cursorCol--
			}
		case "right":
			if m.cursorCol < len(m.lines[m.cursorRow]) {
				m.cursorCol++
			}
		case "up":
			if m.cursorRow > 0 {
				m.cursorRow--
				// if the line above is shorter, clamp the cursor so it doesn't go out of bounds
				if m.cursorCol > len(m.lines[m.cursorRow]) {
					m.cursorCol = len(m.lines[m.cursorRow])
				}
			}

		case "down":
			if m.cursorRow < len(m.lines)-1 {
				m.cursorRow++
				// Same clamping for the line below
				if m.cursorCol > len(m.lines[m.cursorRow]) {
					m.cursorCol = len(m.lines[m.cursorRow])
				}
			}

		// Regular character typing
		default:
			char := msg.String()
			if len(char) == 1 {
				line := m.lines[m.cursorRow]
				m.lines[m.cursorRow] = line[:m.cursorCol] + char + line[m.cursorCol:]
				m.cursorCol++

			}

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
