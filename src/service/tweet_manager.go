package service

import "github.com/goAcademy/src/domain"
import "fmt"

// var tweet *domain.Tweet
var tweets []*domain.Tweet

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

	tweets = append(tweets, tweet2)
	return nil
}

func GetTweets() []*domain.Tweet {
	return tweets
}
func InitializeService() []*domain.Tweet {
	tweets := make([]*domain.Tweet, 4)
	return tweets
}
