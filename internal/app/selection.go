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

func (m model) selectionBounds() (startRow, startCol, endRow, endCol int) {
	// cursor is before select start
	if m.cursorRow < m.selectStartRow || (m.cursorRow == m.selectStartRow && m.cursorCol < m.selectStartCol) {
		return m.cursorRow, m.cursorCol, m.selectStartRow, m.selectStartCol
	}

	return m.selectStartRow, m.selectStartCol, m.cursorRow, m.cursorCol
}

func (m model) isSelected(row, col int) bool {
	if !m.selecting {
		return false
	}

	startRow, startCol, endRow, endCol := m.selectionBounds()

	// is this position between start and end
	if row < startRow || row > endRow {
		return false
	}

	if row == startRow && col < startCol {
		return false
	}

	if row == endRow && col >= endCol {
		return false
	}

	return true
}
