package domain

import "time"

type Tweet struct {
	User string
	Text string
	Date *time.Time
}

func NewTweet(userIn, textIn string) *Tweet {
	time2 := time.Now()
	tweet := Tweet{userIn, textIn, &time2}
	return &tweet
}
