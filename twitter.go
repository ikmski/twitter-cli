package main

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
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

	fmt.Printf("%s %s (%s)\n", t.User.Name, t.User.ScreenName, t.CreatedAt)
	fmt.Printf("%s\n", t.FullText)
	fmt.Printf("%s\n\n", "------------------------------------------------")

}
