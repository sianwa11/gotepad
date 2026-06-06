package app

import (
	tea "charm.land/bubbletea/v2"
)

type buffer struct {
	// Content
	content string
	lines   []string

	// Cursor Position
	cursorRow int
	cursorCol int

	// Viewport (for scrolling)
	offsetRow int // first visible line (for vertical scrolling)
	offsetCol int

	// File metadata
	filename string
	saved    bool

	// Copy/Paste
	selecting      bool
	selectStartRow int
	selectStartCol int
}

type model struct {
	// Content
	content string
	lines   []string

	// Cursor Position
	cursorRow int
	cursorCol int

	// Viewport (for scrolling)
	offsetRow  int // first visible line (for vertical scrolling)
	offsetCol  int
	viewWidth  int // terminal width
	viewHeight int // terminal height

	// File metadata
	filename string
	saved    bool

	// tabs
	buffers      []buffer
	activeBuffer int

	// prompt state
	inputMode     bool
	promptKind    string
	promptValue   string
	statusMessage string

	// Copy/Paste
	selecting      bool
	selectStartRow int
	selectStartCol int
	clipboard      string

	// Debug
	lastClickX int
	lastClickY int
}

func cloneLines(src []string) []string {
	dst := make([]string, len(src))
	copy(dst, src)
	return dst
}

func (m *model) loadActiveBuffer() {
	if len(m.buffers) == 0 {
		m.buffers = []buffer{
			{
				lines: []string{""},
				saved: true,
			},
		}
		m.activeBuffer = 0
	}

	if m.activeBuffer < 0 {
		m.activeBuffer = 0
	}
	if m.activeBuffer >= len(m.buffers) {
		m.activeBuffer = len(m.buffers) - 1
	}

	b := m.buffers[m.activeBuffer]

	m.content = b.content
	m.lines = cloneLines(b.lines)

	m.cursorRow = b.cursorRow
	m.cursorCol = b.cursorCol

	m.offsetRow = b.offsetRow
	m.offsetCol = b.offsetCol

	m.filename = b.filename
	m.saved = b.saved

	m.selecting = b.selecting
	m.selectStartRow = b.selectStartRow
	m.selectStartCol = b.selectStartCol
}

func (m *model) saveActiveBuffer() {
	if len(m.buffers) == 0 || m.activeBuffer < 0 || m.activeBuffer >= len(m.buffers) {
		return
	}

	m.buffers[m.activeBuffer].content = m.content
	m.buffers[m.activeBuffer].lines = cloneLines(m.lines)

	m.buffers[m.activeBuffer].cursorRow = m.cursorRow
	m.buffers[m.activeBuffer].cursorCol = m.cursorCol

	m.buffers[m.activeBuffer].offsetRow = m.offsetRow
	m.buffers[m.activeBuffer].offsetCol = m.offsetCol

	m.buffers[m.activeBuffer].filename = m.filename
	m.buffers[m.activeBuffer].saved = m.saved

	m.buffers[m.activeBuffer].selecting = m.selecting
	m.buffers[m.activeBuffer].selectStartRow = m.selectStartRow
	m.buffers[m.activeBuffer].selectStartCol = m.selectStartCol
}

// func saveActiveBuffer(m model) model {}

func InitialModel() model {
	m := model{
		viewWidth:    0,
		viewHeight:   0,
		clipboard:    "",
		buffers:      []buffer{{lines: []string{""}, saved: true}},
		activeBuffer: 0,
	}
	m.loadActiveBuffer()
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}
