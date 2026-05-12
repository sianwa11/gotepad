package app

import (
	"os"
	"strings"
)

func (m model) handleSave() model {
	if m.filename == "" {
		m.filename = "untitled.txt"
	}

	// save to disk
	file := strings.Join(m.lines, "\n")

	err := os.WriteFile(m.filename, []byte(file), 0644)

	if err != nil {
		m.saved = false
	} else {
		m.saved = true
	}
	return m
}
