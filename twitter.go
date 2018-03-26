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

func (t *twitter) newApi() *anaconda.TwitterApi {

	return anaconda.NewTwitterApiWithCredentials(
		config.AccessToken,
		config.AccessTokenSecret,
		config.ConsumerKey,
		config.ConsumerSecret)

}

func (t *twitter) userStream() {

	api := t.newApi()
	defer api.Close()

	tweets, _ := api.GetHomeTimeline(url.Values{})
	for i := len(tweets) - 1; i >= 0; i-- {
		output(tweets[i])
	}

	stream := api.UserStream(url.Values{})

	for content := range stream.C {

		switch tweet := content.(type) {

		case anaconda.Tweet:

			output(tweet)
		}

	}

}

func (t *twitter) publicStream(q string) {

	api := t.newApi()
	defer api.Close()

	res, _ := api.GetSearch(q, nil)
	for i := len(res.Statuses) - 1; i >= 0; i-- {
		output(res.Statuses[i])
	}

	v := url.Values{}
	v.Set("track", q)
	stream := api.PublicStreamFilter(v)

	for content := range stream.C {

		switch tweet := content.(type) {

		case anaconda.Tweet:

			output(tweet)
		}

	}

}

func output(t anaconda.Tweet) {

	date, err := t.CreatedAtTime()
	if err != nil {
		return
	}

	color.New(color.Reset).Printf("%s\n", "---")
	color.New(color.FgBlue).Printf("%s @%s  ", t.User.Name, t.User.ScreenName)
	color.New(color.FgMagenta).Printf("(%s)\n", date.Local().Format("01/02 15:04:05"))
	color.New(color.Reset).Printf("%s\n\n", t.FullText)

}
