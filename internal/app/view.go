package app

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {

	var sb strings.Builder

	for rowIdx, line := range m.lines {
		if rowIdx == m.cursorRow {
			// this is the line the cursor is on
			// we render it in three parts: before cursor, the cursor block, after cursor
			before := line[:m.cursorCol]
			after := line[m.cursorCol:]

			// If the cursos is at the end of the line, show it as a block on empty space
			cursorChar := " "
			if len(after) > 0 {
				cursorChar = string(after[0])
				after = after[1:]
			}

			sb.WriteString(before + fmt.Sprintf("\033[7m%s\033[0m", cursorChar) + after)
		}else {
			sb.WriteString(line)
		}
		sb.WriteString("\n")
		
	}

		// Simple status line at the bottom
	sb.WriteString(fmt.Sprintf("\n-- Ln %d, Col %d --   ctrl+c to quit",
		m.cursorRow+1, m.cursorCol+1))

		return tea.NewView(sb.String())
}