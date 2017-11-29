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
/*
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
}*/

func isValidTweet(t *testing.T, firstPublishedTweet *domain.Tweet, user string, text string, id int) bool {
	res := firstPublishedTweet.User == user && firstPublishedTweet.Text == text //&& service.GetTweets[id] == firstPublishedTweet
	return res
}

/*
func TestCanRetrieveTweetById(t *testing.T) {

	service.InitializeService()

	var tweet *domain.Tweet
	var id int

	var user domain.User
	user.User = "grupoesfera"
	text := "tweet"

	tweet = domain.NewTweet(user, text)

	id, _ = service.PublishTweet(tweet)

	publishedTweet := service.GetTweetById(id)

	isValidTweet(t, publishedTweet, user, text, id)
}*/

func TestMap(t *testing.T) {
	var tweet, secondTweet, thirdTweet *domain.Tweet

	var twitter service.Twitter

	twitter.InitializeService()
	userIn := "grupoesfera"
	user := domain.NewUser(userIn, "", "", "")
	anotherUserIn := "nick"
	anotherUser := domain.NewUser(anotherUserIn, "", "", "")
	text := "Twiit"
	secondText := "2do Twiit"
	tweet = domain.NewTweet(user.User, text)
	secondTweet = domain.NewTweet(user.User, secondText)
	thirdTweet = domain.NewTweet(anotherUser.User, text)

	twitter.Register(user)
	twitter.Register(anotherUser)
	twitter.LogIn("", "")
	twitter.LogIn("", "")

	firstId, _ := twitter.PublishTweet(tweet)
	twitter.PublishTweet(secondTweet)
	twitter.PublishTweet(thirdTweet)

	tweets := twitter.GetTweetsByUser(user.User)

	firstPublishedTweet := tweets[0]
	//secondPublishedTweet := tweets[1]

	if !isValidTweet(t, firstPublishedTweet, user.User, text, firstId) {
		return
	}
	// if !isValidTweet(t, secondPublishedTweet, user.User, text, secondId) {
	// 	return
	// }
	// if !isValidTweet(t, secondPublishedTweet, user.User, text, thirdId) {
	// 	return
	// }

	twitter.Erase(text, userIn)
	twitter.Erase(secondText, userIn)

	tweetsUser := twitter.GetTweetsByUser(userIn)

	println(tweetsUser)

	if len(twitter.GetTweetsByUser(userIn)) != 0 {
		t.Errorf("expected size 0 but was %d", len(twitter.GetTweetsByUser(userIn)))
	}

}
