package main

import (
	"strconv"

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

			userIn := c.ReadLine()

			user := domain.NewUser(userIn)

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

			c.Print("Write the tweet id: ")

			id := c.ReadLine()
			idAux, _ := strconv.Atoi(id)

			tweets := service.GetTweets()

			c.Println(tweets[idAux].Text)
			c.Println("@", tweets[idAux].User)
			c.Println(tweets[idAux].Date)

			return
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := service.GetTweets()

			for i := 0; i < len(tweets); i++ {
				c.Println(tweets[i].Text)
				c.Println("@", tweets[i].User)
				c.Println(tweets[i].Date)
			}
			return
		},
	})

	// shell.AddCmd(&ishell.Cmd{
	// 	Name: "showUser",
	// 	Help: "Shows a user",
	// 	Func: func(c *ishell.Context) {

	// 		defer c.ShowPrompt(true)

	// 		tweet := service.GetTweet()

	// 		c.Println(tweet.User)

	// 		return
	// 	},
	// })

	shell.Run()

}
