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

	tabBarStyle = lipgloss.NewStyle().Background(lipgloss.Color("236"))
	tabActive   = lipgloss.NewStyle().Background(lipgloss.Color("33")).Foreground(lipgloss.Color("255")).Padding(0, 1)
	tabInactive = lipgloss.NewStyle().Background(lipgloss.Color("240")).Foreground(lipgloss.Color("252")).Padding(0, 1)
	helpStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("246"))
)

func (m model) View() tea.View {
	if m.viewHeight == 0 || m.viewWidth == 0 {
		return tea.NewView("window size not ready yet")
	}

	var sb strings.Builder

	// Top Tabs
	sb.WriteString(m.renderTabBar())
	sb.WriteString("\n")

	// Editor body
	// reverse: 1 row tab bar + 1 row status/help
	maxRows := m.viewHeight - 3
	rendered := 0
	for row := m.offsetRow; row < len(m.lines) && rendered < maxRows; row++ {
		rendered += m.renderLine(&sb, row, maxRows-rendered)
	}

	// fill unused rows to keep footer pinned
	for rendered < maxRows {
		sb.WriteString("\n")
		rendered++
	}

	// Bottom status/help
	sb.WriteString(m.renderStatusBar())

	v := tea.NewView(sb.String())
	v.MouseMode = tea.MouseModeCellMotion
	return v
}

func (m model) renderTabBar() string {
	if len(m.buffers) == 0 {
		return tabBarStyle.Width(m.viewWidth).Render("")
	}

	tabs := make([]string, 0, len(m.buffers))
	for i, b := range m.buffers {
		name := b.filename
		if name == "" {
			name = fmt.Sprintf("Untitled %d", i+1)
		}

		dirty := ""
		if !b.saved {
			dirty = " *"
		}

		label := fmt.Sprintf("%d:%s%s", i+1, name, dirty)

		if i == m.activeBuffer {
			tabs = append(tabs, tabActive.Render(label))
		} else {
			tabs = append(tabs, tabInactive.Render(label))
		}
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, tabs...)
	return tabBarStyle.Width(m.viewWidth).Render(row)
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

	dirtyText := ""
	if !m.saved {
		dirtyText = " [modified]"
	}

	left := statusStyle.Padding(0, 1).Render(
		fmt.Sprintf("Tab %d/%d  %s%s", m.activeBuffer+1, len(m.buffers), filename, dirtyText),
	)

	rightText := fmt.Sprintf("Ln %d, Col %d  %d chars", m.cursorRow+1, m.cursorCol+1, charCount)
	right := statusStyle.Padding(0, 1).Render(rightText)

	w := lipgloss.Width
	midWidth := m.viewWidth - w(left) - w(right)
	if midWidth < 0 {
		midWidth = 0
	}

	help := helpStyle.Render("Tabs: Ctrl+N new  Alt+H prev  Alt+L next  F6/F7 also works")
	mid := statusStyle.Width(midWidth).Render(help)

	return lipgloss.JoinHorizontal(lipgloss.Top, left, mid, right)
}
