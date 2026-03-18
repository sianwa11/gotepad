package app

import tea "charm.land/bubbletea/v2"

func (m model) handleEditing(msg tea.KeyPressMsg) model {

	switch msg.String() {
	case "enter":
		currentLine := m.lines[m.cursorRow]
		before := currentLine[:m.cursorCol]
		after := currentLine[m.cursorCol:]

		// replace the current line with the left part
		m.lines[m.cursorRow] = before
		newLines := make([]string, len(m.lines)+1)
		copy(newLines, m.lines[:m.cursorRow+1])
		newLines[m.cursorRow+1] = after
		copy(newLines[m.cursorRow+2:], m.lines[m.cursorRow+1:])
		m.lines = newLines

		// move cursor to the start of the new line
		m.cursorRow++
		m.cursorCol = 0

	case "backspace":
		if m.cursorCol > 0 {
			line := m.lines[m.cursorRow]
			m.lines[m.cursorRow] = line[:m.cursorCol-1] + line[m.cursorCol:]
			m.cursorCol--
		} else if m.cursorRow > 0 {
			// cursor is at the start of the line - merge this line with the one above
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

	default:
		char := msg.String()
		if len(char) == 1 {
			line := m.lines[m.cursorRow]
			m.lines[m.cursorRow] = line[:m.cursorCol] + char + line[m.cursorCol:]
			m.cursorCol++
		}
	}

	return m
}
