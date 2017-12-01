package domain

import (
	"time"
	//"fmt"
)
type Tweet interface {
	PrintableTweet() string
	GetUser() User
	GetText() string
	SetText(string string)
}


type TextTweet struct {
	User User
	Text string
	Date *time.Time
}

func (t* TextTweet) PrintableTweet() string {
	res := "@"
	res += t.User.User
	res += ": "
	res += t.Text

	return res
}

func (t *TextTweet)GetText() string{
	return t.Text
}

func (t *TextTweet)SetText(text string) {
	t.Text = text
}

func (t *TextTweet)String () string{
	return t.PrintableTweet()
}

func (t *TextTweet)GetUser() User{
	return t.User
}

type ImageTweet struct {
	TextTweet
	link string
}

func (t *ImageTweet)String () string{
	return t.PrintableTweet()
}

func (t *ImageTweet)GetText() string{
	return t.TextTweet.GetText()
}

func (t *ImageTweet)SetText(text string) {
	t.Text = text
}

func (t *ImageTweet)GetUser() User{
	return t.User
}

func (t* ImageTweet) PrintableTweet() string {
	res := t.TextTweet.PrintableTweet()
	res += " "
	res += t.link

	return res
}

func (t *QuoteTweet)String () string{
	return t.PrintableTweet()
}

func (t *QuoteTweet)GetUser() User{
	return t.User
}

func (t *QuoteTweet)GetText() string{
	return t.TextTweet.GetText()
}

func (t *QuoteTweet)SetText(text string) {
	t.Text = text
}

type QuoteTweet struct {
	TextTweet
	Mension Tweet
}

func (t* QuoteTweet) PrintableTweet() string {
	res := t.TextTweet.PrintableTweet()
	res += " "
	res += "\""
	res += t.Mension.PrintableTweet()
	res += "\""
	return res
}



func NewTextTweet(userIn User, textIn string) *TextTweet {
	time2 := time.Now()
	tweet := TextTweet{userIn, textIn, &time2}
	return &tweet
}

func NewImageTweet(userIn User, textIn, link string) *ImageTweet {
	time2 := time.Now()
	tweetAux := TextTweet{userIn, textIn, &time2}
	tweet := ImageTweet{tweetAux, link}
	return &tweet
}

func NewQuoteTweet(userIn User, textIn string, quoted Tweet) *QuoteTweet {
	time2 := time.Now()
	tweetAux := TextTweet{userIn, textIn, &time2}
	tweet := QuoteTweet{tweetAux, quoted}
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