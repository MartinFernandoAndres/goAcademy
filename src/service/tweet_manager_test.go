package service_test

import (
	"testing"

	"github.com/goAcademy/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// var tweet string = "This is my first tweet"

	tweet := "This is my first tweet"

	service.PublishTweet(tweet)

	if service.GetTweet() != tweet {
		t.Error("Expected tweet us ", tweet)
	}

}
