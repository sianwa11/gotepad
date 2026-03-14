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

			// is the cursor inside this chunk?
			if actualRow == m.cursorRow && m.cursorCol >= chunkStart && m.cursorCol <= chunkEnd {

				// find cursor position within just this chunk
				localCol := m.cursorCol - chunkStart

				before := chunk[:localCol]
				after := chunk[localCol:]

				// what character is the cursor sitting on?
				cursorChar := " " // default to space if cursor is at end of line
				if len(after) > 0 {
					cursorChar = string(after[0])
					after = after[1:] // remove it from after since we're drawing it separately
				}

				sb.WriteString(before + "\033[7m" + cursorChar + "\033[0m" + after)

			} else {
				// cursor not in this chunk, draw normally
				sb.WriteString(chunk)
			}

			sb.WriteString("\n")
			rendered++
		}
	}

	sb.WriteString(fmt.Sprintf("\n-- Ln %d, Col %d --   ctrl+c to quit",
		m.cursorRow+1, m.cursorCol+1))

	v := tea.NewView(sb.String())
	v.MouseMode = tea.MouseModeCellMotion
	return v
}
