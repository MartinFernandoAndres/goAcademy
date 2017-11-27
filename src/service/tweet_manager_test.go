package service_test

import (
	"testing"

	"github.com/goAcademy/src/domain"
	"github.com/goAcademy/src/service"
)

/*
func TestPublishedTweetIsSaved(t *testing.T) {
	var tweet *domain.Tweet

	user := "grupoEsfera"
	text := "this is my first tweet"

	tweet = domain.NewTweet(user, text)

	service.PublishTweet(tweet)

	publishedTweet := service.GetTweet()

	if publishedTweet.User != user &&
		publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, publishedTweet.User, publishedTweet.Text)
	}

	if publishedTweet.Date == nil {
		t.Error("expecterd dat can't be nil")
	}

}*/

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	var tweet *domain.Tweet

	user := "grupoEsfera"
	var text string
	text = ""

	tweet = domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestAsdf(t *testing.T) {
	service.InitializeService()

	var tweet, secondTweet *domain.Tweet

	user := "grupoesfera"
	text := "tweet"
	secondText := "2do tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)

	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)

	publishedTweet := service.GetTweets()

	if len(publishedTweet) != 2 {
		t.Errorf("Expect size is 2 but was %d", len(publishedTweet))
		return
	}

	firstPublishedTweet := publishedTweet[0]
	secondPublishedTweet := publishedTweet[1]

	if !isValidTweet(t, firstPublishedTweet, user, text) {
		return
	}
	if !isValidTweet(t, secondPublishedTweet, user, text) {
		return
	}
}

func isValidTweet(t *testing.T, firstPublishedTweet *domain.Tweet, user string, text string) bool {
	res := firstPublishedTweet.User == user && firstPublishedTweet.Text == text
	return res
}
