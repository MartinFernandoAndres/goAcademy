package domain

type Tweet struct {
	User string
	Text string
	//Date time.Date
}

func NewTweet(userIn, textIn string) *Tweet {
	tweet := Tweet{userIn, textIn /*, time.Date.Now()*/}
	return &tweet
}
