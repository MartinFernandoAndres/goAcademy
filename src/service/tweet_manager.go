package service

import "github.com/goAcademy/src/domain"

var tweet *domain.Tweet

func PublishTweet(tweet2 *domain.Tweet) {
	tweet = tweet2
}

func GetTweet() *domain.Tweet {
	return tweet
}
