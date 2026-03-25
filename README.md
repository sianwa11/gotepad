# gotepad

A terminal-based text editor built from scratch in Go using the [Bubble Tea v2](https://charm.land/bubbletea/v2) TUI framework. Built as a deep dive into Go, TUI architecture, and the data structures behind text editors.

> 🚧 **Work in progress** — core editing features are implemented, more on the way.

---

## Demo

```
┌─────────────────────────────────────┐
│ Hello World                         │
│ This is gotepad running in          │
│ your terminal.                      │
│ Click anywhere to place the cu█sor  │
│                                     │
│                                     │
│ -- Ln 4, Col 18 | ctrl+c to quit -- │
└─────────────────────────────────────┘
```

---

## Features

### ✅ Implemented

- **Multi-line editing** — type, delete, and insert text across multiple lines
- **Enter / Backspace** — split lines on enter, merge lines on backspace at line start
- **Arrow key navigation** — move cursor in all directions with proper wrap-aware up/down
- **Wrap-aware movement** — up/down navigates screen rows, not just document lines, so long wrapped lines feel natural
- **Viewport scrolling** — documents longer than the terminal height scroll automatically as you type
- **Line wrapping** — lines wider than the terminal wrap visually without modifying the underlying data
- **Click to place cursor** — click anywhere on screen to instantly move the cursor there, including on wrapped lines
- **Home / End keys** — jump to the start or end of a line instantly
- **Text selection** — hold shift and use arrow keys to select text, with visual blue highlighting

### 🚧 In Progress

- Copy / paste
- Save to file
- Open files
- File menu

### 📋 Planned

- Search and replace
- Line numbers
- Syntax highlighting

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
| `Ctrl+C` | Quit |

---

## Project Structure

```
gotepad/
├── main.go
└── internal/
    └── app/
        ├── model.go       # data structures
        ├── update.go      # main event loop
        ├── view.go        # terminal rendering
        ├── movement.go    # cursor movement handlers
        ├── editing.go     # typing, enter, backspace
        ├── selection.go   # text selection logic
        
```

---

## What I Learned

This project was built incrementally as a way to learn Go. Key concepts explored:

- **Bubble Tea architecture** — the Model/Update/View pattern for TUI apps
- **Text buffer management** — storing document content as `[]string` and manipulating it on keypress
- **Viewport and scrolling** — rendering only the visible portion of a document using an offset
- **Line wrapping** — chopping lines into screen-width chunks visually without modifying underlying data
- **Coordinate mapping** — converting between screen coordinates and document coordinates for mouse clicks
- **Go slice operations** — inserting and deleting from slices without a built-in insert function
- **Pointer vs value receivers** — learned the hard way when `scrollToCursor` wasn't updating state

---

## Built With

- [Go](https://golang.org/)
- [Bubble Tea v2](https://charm.land/bubbletea/v2)

---

*More features coming. Follow progress on [sianwaatemi.com](https://sianwaatemi.com)*