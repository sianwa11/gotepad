package app

import tea "charm.land/bubbletea/v2"

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {

			// Quit
			case "ctrl+c", "q":
				return m, tea.Quit	
				
			// Enter (insert a new line)
			case "enter":
				currentLine := m.lines[m.cursorRow]
				before := currentLine[:m.cursorCol] // everything left of cursor
				after := currentLine[m.cursorCol:] // everything right of cursor

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
				if m.cursorRow < len(m.lines) - 1 {
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

	return m, nil
}