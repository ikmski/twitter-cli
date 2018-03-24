package main

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
	"github.com/mattn/go-runewidth"
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

	for _, r := range runes {
		termbox.SetCell(x, y, r, s.color, s.bgColor)
		x += runewidth.RuneWidth(r)
	}
}

func (s *screen) render(buf *buffer) {

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	list := buf.toList()
	j := 0
	for i := len(list) - 1; i >= 0; i-- {

		var item = list[i]
		if item == nil {
			continue
		}

		v := item.(anaconda.Tweet)

		s.renderLine(0, j*4+0, fmt.Sprintf("%s %s (%s)", v.User.Name, v.User.ScreenName, v.CreatedAt))
		s.renderLine(4, j*4+1, v.FullText)
		s.renderLine(0, j*4+2, "------------------------------------------------")

		j++
	}

	termbox.Flush()
}
