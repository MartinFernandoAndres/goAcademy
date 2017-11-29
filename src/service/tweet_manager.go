package service

import "github.com/goAcademy/src/domain"
import "fmt"

// var tweet *domain.Tweet

type Twitter struct {
	tweets       []*domain.Tweet
	tweetsByUser map[string][]*domain.Tweet
	users        []domain.User
	loguedUsers  []domain.User
}

func (t *Twitter) PublishTweet(tweet2 *domain.Tweet) (int, error) {
	if tweet2.Text == "" {
		return -1, fmt.Errorf("text is required")
	}
	if tweet2.User == "" {
		return -1, fmt.Errorf("user is required")
	}
	if len(tweet2.Text) > 140 {
		return -1, fmt.Errorf("text exceeds 140 characters")
	}

	if !t.isRegistered(tweet2.User) {
		return -1, fmt.Errorf("User is not registed")
	}

	if !t.isLoggued(tweet2.User) {
		return -1, fmt.Errorf("User is not logged")
	}
	if t.isDuplicated(tweet2) {
		return -1, fmt.Errorf("tweet is duplicated")
	}
	t.tweetsByUser[tweet2.User] = append(t.tweetsByUser[tweet2.User], tweet2)
	t.tweets = append(t.tweets, tweet2)

	return len(t.tweets) - 1, fmt.Errorf("tweet ok")
}

func (t *Twitter) Erase(text, user string) {
	for i := 0; i < len(t.tweets); i++ {
		tweetsAux := make([]*domain.Tweet, 0)
		if t.tweets[i].User == user && t.tweets[i].Text == text {
			for j := 0; j < i; j++ {
				tweetsAux = append(tweetsAux, t.tweets[j])
			}
			for j := i; j < len(t.tweets); j++ {
				tweetsAux = append(tweetsAux, t.tweets[j])
			}
		}
		t.tweets = tweetsAux
	}
	for i := 0; i < len(t.GetTweetsByUser(user)); i++ {
		tweetsAux := make([]*domain.Tweet, 0)
		for j := 0; j < len(t.GetTweetsByUser(user)); j++ {
			if t.GetTweetsByUser(user)[j].Text != text {
				tweetsAux = append(tweetsAux, t.GetTweetsByUser(user)[j])
			}
		}
		t.tweetsByUser[user] = tweetsAux
	}
}

func (t *Twitter) isLoggued(user string) bool {
	for i := 0; i < len(t.loguedUsers); i++ {
		if t.loguedUsers[i].User == user {
			return true
		}
	}
	return false

}

func (t *Twitter) isDuplicated(tweet *domain.Tweet) bool {
	for i := 0; i < len(t.GetTweetsByUser(tweet.User)); i++ {
		if t.GetTweetsByUser(tweet.User)[i].Text == tweet.Text {
			return true
		}
	}
	return false
}

func (t *Twitter) LogIn(name, pass string) {
	for i := 0; i < len(t.users); i++ {
		if (t.users[i].Mail == name || t.users[i].Nick == name) && t.users[i].Pass == pass && !t.isLoggued(t.users[i].User) {
			t.loguedUsers = append(t.loguedUsers, t.users[i])
		}
	}
}

func (t *Twitter) GetTweets() []*domain.Tweet {
	return t.tweets
}

func (t *Twitter) GetUsers() []domain.User {
	return t.users
}

func (t *Twitter) isRegistered(user string) bool {
	for i := 0; i < len(t.users); i++ {
		if t.users[i].User == user {
			return true
		}
	}
	return false
}

func (t *Twitter) InitializeService() ([]*domain.Tweet, map[string][]*domain.Tweet, []domain.User) {
	t.tweets = make([]*domain.Tweet, 0)
	t.tweetsByUser = make(map[string][]*domain.Tweet)
	t.users = make([]domain.User, 0) //ACA PUEDE FALLAR EL =s
	return t.tweets, t.tweetsByUser, t.users
}

func (t *Twitter) LogOut(user string) {
	for i := 0; i < len(t.loguedUsers); i++ {
		if t.loguedUsers[i].User == user {
			for j := 0; j < i; j++ {
				t.loguedUsers = append(t.loguedUsers, t.loguedUsers[j])
			}
			for j := i; j < len(t.loguedUsers); j++ {
				t.loguedUsers = append(t.loguedUsers, t.loguedUsers[j])
			}
		}
	}
}

func (t *Twitter) GetTweetById(id int) *domain.Tweet {
	return t.tweets[id]
}

func (t *Twitter) GetUser(user string) domain.User {
	for i := 0; i < len(t.users); i++ {
		if t.users[i].User == user {
			return t.users[i]
		}
	}
	return domain.NewUser("", "", "", "")
}

func (t *Twitter) GetTweetsByUser(user string) []*domain.Tweet {
	return t.tweetsByUser[user]
}

func (t *Twitter) Register(user domain.User) (int, error) {
	if t.isRegistered(user.User) {
		return -1, fmt.Errorf("User is already registed")
	}
	t.users = append(t.users, user)
	return 0, fmt.Errorf("")
}
