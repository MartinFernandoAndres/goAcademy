package domain

import "time"

type Tweet struct {
	User User
	Text string
	Date *time.Time
}

func NewTweet(userIn User, textIn string) *Tweet {
	time2 := time.Now()
	tweet := Tweet{userIn, textIn, &time2}
	return &tweet
}

type User struct {
	User string
}

func NewUser(user string) User {
	User := User{user}
	return User
}
