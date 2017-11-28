package domain

import "time"

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
	User string
	Mail string
	Nick string
	Pass string
}

func NewUser(user, mail, nick, pass string) User {
	User := User{user, mail, nick, pass}
	return User
}
