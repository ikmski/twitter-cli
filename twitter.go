package main

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	termbox "github.com/nsf/termbox-go"
)

type twitter struct {
	screen *screen
	buffer *buffer

	updateCh chan bool
	eventCh  chan termbox.Event
}

func newTwitter() *twitter {

	t := new(twitter)

	t.eventCh = make(chan termbox.Event)
	t.updateCh = make(chan bool)

	t.screen = newScreen()
	t.buffer = newBuffer(32)

	return t
}

func (t *twitter) userStream() {

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	go func() {
		for {
			t.eventCh <- termbox.PollEvent()
		}
	}()

	go func() {

		api := anaconda.NewTwitterApiWithCredentials(
			config.AccessToken,
			config.AccessTokenSecret,
			config.ConsumerKey,
			config.ConsumerSecret)

		list, _ := api.GetHomeTimeline(url.Values{})
		for i := len(list) - 1; i >= 0; i-- {
			t.buffer.push(list[i])
		}
		t.updateCh <- true

		stream := api.UserStream(url.Values{})

		for tweet := range stream.C {

			switch content := tweet.(type) {

			case anaconda.Tweet:

				t.buffer.push(content)
				t.updateCh <- true

			}

		}

	}()

	t.screen.render(t.buffer)

mainloop:
	for {

		select {
		case e := <-t.eventCh:

			switch e.Type {

			case termbox.EventKey:

				switch e.Key {

				case termbox.KeyEsc:
					break mainloop

				default:

				}

			default:

			}

		case <-t.updateCh:
			t.screen.render(t.buffer)

		}
	}

}
