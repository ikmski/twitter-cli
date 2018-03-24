package main

import (
	"time"

	termbox "github.com/nsf/termbox-go"
)

type screen struct {
	color   termbox.Attribute
	bgColor termbox.Attribute
}

func newScreen() *screen {

	s := new(screen)

	s.color = termbox.ColorDefault
	s.bgColor = termbox.ColorDefault

	return s
}

func (s *screen) renderLine(x, y int, str string) {

	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		termbox.SetCell(x+i, y, runes[i], s.color, s.bgColor)
	}
}

func (s *screen) render(buf *buffer) {

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	s.renderLine(0, 0, "Press ESC to exit.")
	s.renderLine(0, 0, time.Now().Format(time.UnixDate))

	termbox.Flush()
}
