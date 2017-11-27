package service_test

import (
	"testing"

	"github.com/goAcademy/src/domain"
	"github.com/goAcademy/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	var tweet *domain.Tweet

	user := "grupoEsfera"
	text := "this is my first tweet"

	tweet = domain.NewTweet(user, text)

	service.PublishTweet(tweet)

	publishedTweet := service.GetTweet()

	// if time.Now().Equal {
	// 	println("asd")
	// }

	if publishedTweet.User != user &&
		publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, publishedTweet.User, publishedTweet.Text)
	}

	// if publishedTweet.Date == nil {
	// 	t.Error("expecterd dat can't be nil")
	// }

}

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
