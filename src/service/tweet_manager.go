package service

import "github.com/goAcademy/src/domain"
import "fmt"

var tweet *domain.Tweet

func PublishTweet(tweet2 *domain.Tweet) error {
	if tweet2.Text == "" {
		return fmt.Errorf("text is required")
	}
	if tweet2.User == "" {
		return fmt.Errorf("user is required")
	}
	if len(tweet2.Text) > 140 {
		return fmt.Errorf("text exceeds 140 characters")
	}

	tweet = tweet2
	return nil
}

func GetTweet() *domain.Tweet {
	return tweet
}
