package app

import tea "charm.land/bubbletea/v2"

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:
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
		}
	}

	return m, nil
}