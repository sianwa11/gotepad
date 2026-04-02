package app

func (m model) handleSelection(key string) model {
	// Start selection once; keep extending on subsequent Shift+arrow presses.
	if !m.selecting {
		m.selectStartRow = m.cursorRow
		m.selectStartCol = m.cursorCol
		m.selecting = true
	}

	// Reuse normal movement logic so wrapped-line behavior matches regular arrows.
	moveKey := normalizeShiftKey(key)
	if moveKey != "" {
		m = m.handleMovement(moveKey)
		m.selecting = true
	}

	return m
}

func normalizeShiftKey(key string) string {
	switch key {
	case "shift+left":
		return "left"
	case "shift+right":
		return "right"
	case "shift+up":
		return "up"
	case "shift+down":
		return "down"
	default:
		return ""
	}
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
