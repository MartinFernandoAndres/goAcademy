package service

import "github.com/goAcademy/src/domain"
import "fmt"

// var tweet *domain.Tweet
var tweets []*domain.Tweet
var tweetsByUser map[domain.User][]*domain.Tweet
var users []domain.User
var userNow domain.User

func PublishTweet(tweet2 *domain.Tweet) (int, error) {
	if tweet2.Text == "" {
		return -1, fmt.Errorf("text is required")
	}
	if tweet2.User.User == "" {
		return -1, fmt.Errorf("user is required")
	}
	if len(tweet2.Text) > 140 {
		return -1, fmt.Errorf("text exceeds 140 characters")
	}

	if !isRegistered(tweet2.User) {
		users = append(users, tweet2.User)
	}

	userNow = tweet2.User

	tweetsByUser[tweet2.User] = append(tweetsByUser[tweet2.User], tweet2)
	tweets = append(tweets, tweet2)

	return len(tweets) - 1, fmt.Errorf("tweet ok")
}

func GetTweets() []*domain.Tweet {
	return tweets
}

func GetUsers() []domain.User {
	return users
}

func isRegistered(user domain.User) bool {
	for i := 0; i < len(users); i++ {
		if users[i].User == user.User {
			return true
		}
	}
	return false
}

func InitializeService() ([]*domain.Tweet, map[domain.User][]*domain.Tweet) {
	tweets := make([]*domain.Tweet, 4)
	tweetsByUser = make(map[domain.User][]*domain.Tweet)
	return tweets, tweetsByUser
}

func GetTweetById(id int) *domain.Tweet {
	return tweets[id]
}

func GetTweetsByUser(user domain.User) []*domain.Tweet {
	return tweetsByUser[user]
}
