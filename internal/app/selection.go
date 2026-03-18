package app

func (m model) handleSelection(key string) model {

	switch key {
	case "shift+left":
		if !m.selecting {
			m.selectStartRow = m.cursorRow
			m.selectStartCol = m.cursorCol
			m.selecting = true
		}
		// move left
		if m.cursorCol > 0 {
			m.cursorCol--
		}

	case "shift+right":
		if !m.selecting {
			m.selectStartRow = m.cursorRow
			m.selectStartCol = m.cursorCol
			m.selecting = true
		}
		// move right
		if m.cursorCol < len(m.lines[m.cursorRow]) {
			m.cursorCol++
		}

	case "shift+up":
		if !m.selecting {
			m.selectStartRow = m.cursorRow
			m.selectStartCol = m.cursorCol
			m.selecting = true
			// move up
			if m.cursorRow > 0 {
				m.cursorRow--
				if m.cursorCol > len(m.lines[m.cursorRow]) {
					m.cursorCol = len(m.lines[m.cursorRow])
				}
			}
		}

	case "shift+down":
		if !m.selecting {
			m.selectStartRow = m.cursorRow
			m.selectStartCol = m.cursorCol
			m.selecting = true
			// move down
			if m.cursorRow < len(m.lines)-1 {
				m.cursorRow++
				if m.cursorCol > len(m.lines[m.cursorRow]) {
					m.cursorCol = len(m.lines[m.cursorRow])
				}
			}
		}

	}
	return m
}
