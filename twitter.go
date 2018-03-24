package main

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

type twitter struct {
	screen *screen
	buffer *buffer

	updateCh chan bool
}

func newTwitter() *twitter {

	t := new(twitter)

	t.updateCh = make(chan bool)

	t.screen = newScreen()
	t.buffer = newBuffer(10)

	return t
}

func (t *twitter) userStream() {

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

	for {

		select {

		case <-t.updateCh:
			t.screen.render(t.buffer.toList())

		}
	}

}
