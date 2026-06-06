# gotepad

A terminal-based text editor built from scratch in Go using the [Bubble Tea v2](https://charm.land/bubbletea/v2) TUI framework. Built as a deep dive into Go, TUI architecture, and the data structures behind text editors.

> 🚧 **Work in progress** - core editing features are implemented, more on the way.

---

## Demo

```
┌─────────────────────────────────────────────────────┐
│ 1:untitled.txt   2:notes.txt *   3:Untitled 3       │
├─────────────────────────────────────────────────────┤
│ Hello World                                         │
│ This is gotepad running in                          │
│ your terminal.                                      │
│ Click anywhere to place the cu█sor                  │
│                                                     │
├─────────────────────────────────────────────────────┤
│ Tab 2/3  notes.txt [modified]    Ln 4, Col 18       │
│ Tabs: Ctrl+N new  Alt+H prev  Alt+L next Ctrl+S save│
└─────────────────────────────────────────────────────┘
```

---

## Features

### ✅ Implemented

- **Multi-line editing** - type, delete, and insert text across multiple lines
- **Enter / Backspace** - split lines on enter, merge lines on backspace at line start
- **Arrow key navigation** - move cursor in all directions with proper wrap-aware up/down
- **Wrap-aware movement** - up/down navigates screen rows, not just document lines
- **Viewport scrolling** - documents longer than the terminal height scroll automatically
- **Line wrapping** - lines wider than the terminal wrap visually without modifying underlying data
- **Click to place cursor** - click anywhere on screen to move the cursor, including on wrapped lines
- **Home / End keys** - jump to the start or end of a line instantly
- **Text selection** - hold Shift and use arrow keys to select text, with visual blue highlighting
- **Copy / Paste** - Ctrl+C to copy selection, Alt+V to paste
- **Save to file** - Ctrl+S saves the active tab; untitled tabs get unique names (untitled.txt, untitled-2.txt, etc.)
- **Multi-tab editing** - open multiple files in tabs, each with independent cursor, scroll, and selection state
- **Tab switching** - Alt+L / F7 next tab, Alt+H / F6 previous tab, Ctrl+N new tab
- **Unsaved indicators** - active tab shows `[modified]` in status bar; tab label shows `*` when unsaved
- **Per-tab save** - each tab saves independently to its own file
- **Tab bar UI** - coloured tab bar at the top with active tab highlighted

### 🚧 In Progress

- **Open file from disk** - Alt+O path prompt to load a file into a new tab
- **Save As** - save active tab's content under a new filename
- **Close tab** - Ctrl+W with unsaved-change protection

### 📋 Planned

- Search and replace
- Line numbers
- Syntax highlighting
- File picker / directory browser

---

## Installation

```bash
git clone https://github.com/sianwa11/gotepad
cd gotepad
go run .
```

**Requirements:** Go 1.21+

---

## Usage

| Key | Action |
|-----|--------|
| Type normally | Insert characters |
| `Enter` | New line |
| `Backspace` | Delete character / merge lines |
| `Arrow keys` | Move cursor |
| `Shift + Arrow` | Select text |
| `Home` | Jump to start of line |
| `End` | Jump to end of line |
| `Click` | Place cursor at click position |
| `Ctrl+C` | Copy selection |
| `Alt+V` | Paste |
| `Ctrl+S` | Save active tab |
| `Ctrl+N` | New tab |
| `Alt+L` / `F7` | Next tab |
| `Alt+H` / `F6` | Previous tab |
| `Esc` | Quit |

---

## Project Structure

```
gotepad/
├── main.go
└── internal/
    └── app/
        ├── model.go       # data structures, buffer type, tab sync
        ├── update.go      # main event loop, tab and prompt routing
        ├── view.go        # terminal rendering, tab bar, status bar
        ├── movement.go    # cursor movement handlers
        ├── editing.go     # typing, enter, backspace
        ├── selection.go   # text selection logic
        ├── copy_paste.go  # copy, paste, delete selection
        └── save.go        # save to disk, unique untitled naming
```

---

## What I Learned

- **Bubble Tea architecture** - the Model/Update/View pattern for TUI apps
- **Text buffer management** - storing document content as `[]string` and manipulating it on keypress
- **Multi-buffer design** - separating per-tab state (`buffer`) from global state (`model`) with load/save sync functions
- **Viewport and scrolling** - rendering only the visible portion of a document using an offset
- **Line wrapping** - chopping lines into screen-width chunks visually without modifying underlying data
- **Coordinate mapping** - converting between screen coordinates and document coordinates for mouse clicks
- **Go slice operations** - inserting and deleting from slices without a built-in insert function
- **Pointer vs value receivers** - learned the hard way when `scrollToCursor` wasn't updating state
- **O(1) dirty tracking** - setting `saved = false` only at mutation points rather than checking on every render

---

## Built With

- [Go](https://golang.org/)
- [Bubble Tea v2](https://charm.land/bubbletea/v2)
- [Lip Gloss v2](https://charm.land/lipgloss/v2)

---

*More features coming. Follow progress on [sianwaatemi.com](https://sianwaatemi.com)*
