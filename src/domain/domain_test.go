package domain_test

import(
	"testing"
	"github.com/goAcademy/src/domain"
)

func TestTextTweetPrintsUserAndText(t *testing.T) {
	userIn := "grupoesfera"
	user := domain.NewUser(userIn, "a", "a", "a")

	// Initialization
	tweet := domain.NewTextTweet(user, "This is my tweet")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {
	userIn := "grupoesfera"
	user := domain.NewUser(userIn, "a", "a", "a")

	// Initializationv
	tweet := domain.NewImageTweet(user, "This is my image", "http://www.grupoesfera.com.ar/common/img/grupoesfera.png")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my image http://www.grupoesfera.com.ar/common/img/grupoesfera.png"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {
	userIn := "grupoesfera"
	user := domain.NewUser(userIn, "a", "a", "a")

	anotherUserIn := "nick"
	anotherUser := domain.NewUser(anotherUserIn, "b", "b", "b")

	// Initialization
	quotedTweet := domain.NewTextTweet(user, "This is my tweet")
	tweet := domain.NewQuoteTweet(anotherUser, "Awesome", quotedTweet)

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := `@nick: Awesome "@grupoesfera: This is my tweet"`
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestCanGetAStringFromATweet(t *testing.T) {
	userIn := "grupoesfera"
	user := domain.NewUser(userIn, "a", "a", "a")

	// Initialization
	tweet := domain.NewTextTweet(user, "This is my tweet")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}
