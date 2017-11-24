package service

var tweet string

func PublishTweet(tweet2 string) {
	tweet = tweet2
}

func GetTweet() string {
	return tweet
}
