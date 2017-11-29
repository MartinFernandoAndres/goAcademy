package domain

import (
	"time"
)

type Tweet struct {
	User string
	Text string
	Date *time.Time
}

func NewTweet(userIn string, textIn string) *Tweet {
	time2 := time.Now()
	tweet := Tweet{userIn, textIn, &time2}
	return &tweet
}

type User struct {
	User             string
	Mail             string
	Nick             string
	Pass             string
	IFollowThisUsers []User
}

func NewUser(user, mail, nick, pass string) User {
	iFollowThisUsers := make([]User, 0)
	User := User{user, mail, nick, pass, iFollowThisUsers}
	return User
}
