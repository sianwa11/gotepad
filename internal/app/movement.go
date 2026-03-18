package app

func (m model) handleMovement(key string) model {

	switch key {
	case "home":
		m.cursorCol = 0

	case "end":
		m.cursorCol = len(m.lines[m.cursorRow])

	case "left":
		m.selecting = false
		if m.cursorCol > 0 {
			m.cursorCol--
		}

	case "right":
		m.selecting = false
		if m.cursorCol < len(m.lines[m.cursorRow]) {
			m.cursorCol++
		}

	case "up":
		m.selecting = false
		currentChunk := m.cursorCol / m.viewWidth
		offsetInChunk := m.cursorCol - (currentChunk * m.viewWidth)

		// is there a chunk above in the same line?
		if currentChunk > 0 {
			prevChunkStart := (currentChunk - 1) * m.viewWidth
			newCol := prevChunkStart + offsetInChunk
			// clamp to end of line
			if newCol > len(m.lines[m.cursorRow]) {
				newCol = len(m.lines[m.cursorRow])
			}
			m.cursorCol = newCol
		} else {
			// move to previous document line
			if m.cursorRow > 0 {
				m.cursorRow--
				prevLine := m.lines[m.cursorRow]

				lastChunkStart := (len(prevLine) / m.viewWidth) * m.viewWidth

				// land on the last chunk at the same horizontal offset
				newCol := lastChunkStart + offsetInChunk

				// clamp to end of line
				if newCol > len(prevLine) {
					newCol = len(prevLine)
				}

				m.cursorCol = newCol
			}
		}

	case "down":
		m.selecting = false
		currentChunk := m.cursorCol / m.viewWidth
		offsetInChunk := m.cursorCol - (currentChunk * m.viewWidth)
		nextChunkStart := (currentChunk + 1) * m.viewWidth

		// does the next chunk exist in the same line?
		if nextChunkStart < len(m.lines[m.cursorRow]) {
			newCol := nextChunkStart + offsetInChunk

			if newCol > len(m.lines[m.cursorRow]) {
				newCol = len(m.lines[m.cursorRow])
			}
			m.cursorCol = newCol
		} else {
			// move to next document line
			if m.cursorRow < len(m.lines)-1 {
				m.cursorRow++
				nextLine := m.lines[m.cursorRow]

				// land at the same horizontal offset within the first chunk
				newCol := offsetInChunk

				if newCol > len(nextLine) {
					newCol = len(nextLine)
				}

				m.cursorCol = newCol
			}

		}

	}

	return m
}
