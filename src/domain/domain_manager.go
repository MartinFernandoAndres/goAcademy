package domain

type Tweet struct {
	User string
	Text string
}

func NewTweet(userIn, textIn string) *Tweet {
	tweet := Tweet{userIn, textIn}
	return &tweet
}
