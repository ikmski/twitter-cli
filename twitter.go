package main

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/fatih/color"
)

type twitter struct {
}

func newTwitter() *twitter {

	t := new(twitter)

	return t
}

func (t *twitter) userStream() {

	api := anaconda.NewTwitterApiWithCredentials(
		config.AccessToken,
		config.AccessTokenSecret,
		config.ConsumerKey,
		config.ConsumerSecret)

	tweets, _ := api.GetHomeTimeline(url.Values{})
	for i := len(tweets) - 1; i >= 0; i-- {
		output(tweets[i])
	}

	stream := api.UserStream(url.Values{})

	for tweet := range stream.C {

		switch content := tweet.(type) {

		case anaconda.Tweet:

			output(content)
		}

	}

}

func output(t anaconda.Tweet) {

	date, err := t.CreatedAtTime()
	if err != nil {
		return
	}

	color.New(color.Reset).Printf("%s\n", "---")
	color.New(color.FgBlue).Printf("%s @%s\n", t.User.Name, t.User.ScreenName)
	color.New(color.Reset).Printf("  %s\n\n", t.FullText)
	color.New(color.FgMagenta).Printf("(%s)\n", date.Local().Format("01/02 15:04:05"))

}
