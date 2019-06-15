package main

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

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
		auth.AccessToken,
		auth.AccessTokenSecret,
		auth.ConsumerKey,
		auth.ConsumerSecret)

}

func (t *twitter) userStream() {

	api := t.newApi()
	defer api.Close()

	tweets, _ := api.GetHomeTimeline(url.Values{})
	for i := len(tweets) - 1; i >= 0; i-- {
		output(tweets[i])
	}

	cursor, err := api.GetFriendsIds(url.Values{})
	friendsIds := cursor.Ids
	if err != nil {
		fmt.Println(err)
		return
	}

	var IDs []string
	for _, id := range friendsIds {
		IDs = append(IDs, strconv.FormatInt(id, 10))
	}

	v := url.Values{}
	v.Set("follow", strings.Join(IDs, ", "))
	stream := api.PublicStreamFilter(v)

	for content := range stream.C {

		switch tweet := content.(type) {

		case anaconda.Tweet:
			if isFriend(tweet.User.Id, friendsIds) {
				output(tweet)
			}
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

func isFriend(id int64, friendsIds []int64) bool {
	for _, v := range friendsIds {
		if id == v {
			return true
		}
	}
	return false
}
