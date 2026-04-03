package app

import "strings"

func (m model) handleCopy() model {
	if !m.selecting {
		return m
	}

	startRow, startCol, endRow, endCol := m.selectionBounds()

	// Same line
	if startRow == endRow {
		m.clipboard = m.lines[startRow][startCol:endCol]
		return m
	}

	// Multiple lines: build piece by piece
	var result string

	// First line
	result += m.lines[startRow][startCol:]

	// Middle lines
	for row := startRow + 1; row < endRow; row++ {
		result += "\n" + m.lines[row]
	}

	// Last line
	result += "\n" + m.lines[endRow][:endCol]

	m.clipboard = result
	return m
}

func (m model) handlePaste() model {
	return m.handlePasteText(m.clipboard)
}

func (m model) handlePasteText(text string) model {
	if text == "" {
		return m
	}

	// If a selection exists, replace it first
	if m.selecting {
		m = m.deleteSelection()
	}

	line := m.lines[m.cursorRow]
	before := line[:m.cursorCol]
	after := line[m.cursorCol:]

	// single-line paste
	if !strings.Contains(text, "\n") {
		m.lines[m.cursorRow] = before + text + after
		m.cursorCol += len(text)
		m.selecting = false
		return m
	}

	// Multi-line paste
	parts := strings.Split(text, "\n")

	first := before + parts[0]
	last := parts[len(parts)-1] + after

	newLines := make([]string, 0, len(m.lines)+len(parts)-1)
	newLines = append(newLines, m.lines[:m.cursorRow]...)
	newLines = append(newLines, first)

	if len(parts) > 2 {
		newLines = append(newLines, parts[1:len(parts)-1]...)
	}

	newLines = append(newLines, last)
	newLines = append(newLines, m.lines[m.cursorRow+1:]...)
	m.lines = newLines

	m.cursorRow += len(parts) - 1
	m.cursorCol = len(parts[len(parts)-1])
	m.selecting = false

	return m
}

func (m model) deleteSelection() model {
	if !m.selecting {
		return m
	}

	startRow, startCol, endRow, endCol := m.selectionBounds()

	// Same-line delete
	if startRow == endRow {
		line := m.lines[startRow]
		m.lines[startRow] = line[:startCol] + line[endCol:]
		m.cursorRow = startRow
		m.cursorCol = startCol
		m.selecting = false
		return m
	}

	// Multi-line delete
	head := m.lines[startRow][:startCol]
	tail := m.lines[endRow][endCol:]

	m.lines[startRow] = head + tail
	m.lines = append(m.lines[:startRow+1], m.lines[endRow+1:]...)

	m.cursorRow = startRow
	m.cursorCol = startCol
	m.selecting = false
	return m
}
