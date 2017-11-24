package domain

import "time"

type Tweet struct {
	User string
	Text string
	Date time.Date
}

func NewTweet(userIn, textIn string) *Tweet {
	tweet := Tweet{userIn, textIn, time.Now().Date}
	return &tweet
}
