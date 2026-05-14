package app

import (
	"fmt"
	"os"
	"strings"
)

func (m model) handleSave() model {
	if m.filename == "" {
		m.filename = m.nextUntitledName()
	}

	// save to disk
	file := strings.Join(m.lines, "\n")
	err := os.WriteFile(m.filename, []byte(file), 0644)

	m.saved = err == nil
	return m
}

func (m model) nextUntitledName() string {
	used := map[string]bool{}
	for _, b := range m.buffers {
		if b.filename != "" {
			used[b.filename] = true
		}
	}

	if !used["untitled.txt"] {
		return "untitled.txt"
	}

	for i := 2; ; i++ {
		name := fmt.Sprintf("untitled-%d.txt", i)
		if !used[name] {
			return name
		}
	}
}
