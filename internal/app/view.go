package app

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

const (
	styleReset    = "\033[0m"
	styleCursor   = "\033[7m"
	styleSelected = "\033[48;5;24m"
)

func (m model) View() tea.View {
	if m.viewHeight == 0 || m.viewWidth == 0 {
		return tea.NewView("window size not ready yet")
	}

	var sb strings.Builder
	maxRows := m.viewHeight - 2
	rendered := 0

	for row := m.offsetRow; row < len(m.lines) && rendered < maxRows; row++ {
		rendered += m.renderLine(&sb, row, maxRows-rendered)
	}

	sb.WriteString(m.renderStatusBar())

	v := tea.NewView(sb.String())
	v.MouseMode = tea.MouseModeCellMotion
	return v
}

func (m model) renderLine(sb *strings.Builder, row int, rowsLeft int) int {
	line := m.lines[row]

	if len(line) == 0 {
		if row == m.cursorRow {
			sb.WriteString(styleCursor + " " + styleReset)
		}
		sb.WriteString("\n")
		return 1
	}

	rendered := 0
	for chunkStart := 0; chunkStart < len(line) && rendered < rowsLeft; chunkStart += m.viewWidth {
		chunkEnd := chunkStart + m.viewWidth
		if chunkEnd > len(line) {
			chunkEnd = len(line)
		}

		m.renderChunk(sb, row, chunkStart, line[chunkStart:chunkEnd])

		if row == m.cursorRow && m.cursorCol == chunkEnd && chunkEnd == len(line) {
			sb.WriteString(styleCursor + " " + styleReset)
		}

		sb.WriteString("\n")
		rendered++
	}

	return rendered
}

func (m model) renderChunk(sb *strings.Builder, row, chunkStart int, chunk string) {
	for i, ch := range chunk {
		col := chunkStart + i
		switch {
		case row == m.cursorRow && col == m.cursorCol:
			sb.WriteString(styleCursor + string(ch) + styleReset)
		case m.isSelected(row, col):
			sb.WriteString(styleSelected + string(ch) + styleReset)
		default:
			sb.WriteString(string(ch))
		}
	}
}

func (m model) renderStatusBar() string {
	return fmt.Sprintf("\n-- Ln %d, Col %d | clip=%d | chunk=%d offsetRow=%d --",
		m.cursorRow+1, m.cursorCol+1, len(m.clipboard), m.cursorCol/m.viewWidth, m.offsetRow)
}
