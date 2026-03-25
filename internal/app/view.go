package app

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	var sb strings.Builder

	rendered := 0

	if m.viewHeight == 0 {
		return tea.NewView("viewHeight is 0! WindowSizeMsg not handled")
	}

	for rowIdx, line := range m.lines[m.offsetRow:] {
		actualRow := rowIdx + m.offsetRow

		if rendered >= m.viewHeight-2 {
			break
		}

		// empty line
		if len(line) == 0 {
			// is the cursor sitting on this empty line?
			if actualRow == m.cursorRow {
				sb.WriteString("\033[7m \033[0m") // draw cursor as highlighted space
			}
			sb.WriteString("\n")
			rendered++
			continue
		}

		// chop line into chunks
		for chunkStart := 0; chunkStart < len(line); chunkStart += m.viewWidth {
			chunkEnd := chunkStart + m.viewWidth
			if chunkEnd > len(line) {
				chunkEnd = len(line)
			}
			chunk := line[chunkStart:chunkEnd]

			// draw character by character
			for i, ch := range chunk {
				actualCol := chunkStart + i

				if actualRow == m.cursorRow && actualCol == m.cursorCol {
					// draw cursor
					sb.WriteString("\033[7m" + string(ch) + "\033[0m")
				} else if m.isSelected(actualRow, actualCol) {
					// draw selected character with blue background
					sb.WriteString("\033[48;5;24m" + string(ch) + "\033[0m")
				} else {
					// draw normally
					sb.WriteString(string(ch))
				}
			}

			// handle cursor at end of line
			if actualRow == m.cursorRow && m.cursorCol == chunkEnd && chunkEnd == len(line) {
				sb.WriteString("\033[7m \033[0m")
			}

			sb.WriteString("\n")
			rendered++
		}
	}

	// sb.WriteString(fmt.Sprintf("\n-- Ln %d, Col %d --   ctrl+c to quit",
	// m.cursorRow+1, m.cursorCol+1))
	// sb.WriteString(fmt.Sprintf("\n-- click=(%d,%d) cursorRow=%d cursorCol=%d offsetRow=%d --",
	// 	m.lastClickX, m.lastClickY, m.cursorRow, m.cursorCol, m.offsetRow))
	fmt.Fprintf(&sb, "\n-- Ln %d, Col %d | chunk=%d offsetRow=%d --",
		m.cursorRow+1, m.cursorCol+1, m.cursorCol/m.viewWidth, m.offsetRow)

	v := tea.NewView(sb.String())
	v.MouseMode = tea.MouseModeCellMotion
	return v
}
