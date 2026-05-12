package app

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	cursorStyle   = lipgloss.NewStyle().Reverse(true)
	selectedStyle = lipgloss.NewStyle().Background(lipgloss.Color("24"))
	statusStyle   = lipgloss.NewStyle().Background(lipgloss.Color("24")).Foreground(lipgloss.Color("255"))
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
			sb.WriteString(cursorStyle.Render(" "))
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
			sb.WriteString(cursorStyle.Render(" "))
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
			sb.WriteString(cursorStyle.Render(string(ch)))
		case m.isSelected(row, col):
			sb.WriteString(selectedStyle.Render(string(ch)))
		default:
			sb.WriteString(string(ch))
		}
	}
}

func (m model) renderStatusBar() string {
	filename := m.filename
	if filename == "" {
		filename = "[No file]"
	}
	charCount := 0
	for _, l := range m.lines {
		charCount += len(l)
	}

	w := lipgloss.Width

	left := statusStyle.Padding(0, 1).Render("  " + filename + "  ")
	right := statusStyle.Padding(0, 1).Render(fmt.Sprintf("Ln %d, Col %d  %d chars", m.cursorRow+1, m.cursorCol+1, charCount))
	mid := statusStyle.Width(m.viewWidth - w(left) - w(right)).Render("")

	return "\n" + lipgloss.JoinHorizontal(lipgloss.Top, left, mid, right)
}
