package main

import (
	"github.com/abiosoft/ishell"
	"github.com/goAcademy/src/domain"
	"github.com/goAcademy/src/service"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			c.Print("Write your user: ")

			user := c.ReadLine()

			tweet := domain.NewTweet(user, text)

			service.PublishTweet(tweet)

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := service.GetTweet()

			c.Println(tweet.Text)
			c.Println("@", tweet.User)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showUser",
		Help: "Shows a user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := service.GetTweet()

			c.Println(tweet.User)

			return
		},
	})

	shell.Run()

}
