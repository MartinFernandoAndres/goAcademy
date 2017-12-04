package main

import (
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/goAcademy/src/domain"
	"github.com/goAcademy/src/service"
	"github.com/goAcademy/src/rest"
)

func main() {


	 twitter := new( service.Twitter)
	twitter.InitializeService()

	ginServer := rest.NewGinServer(twitter)
	ginServer.StartGinServer()

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

			user := twitter.GetUser(userIn)

			tweet := domain.NewTextTweet(user, text)

			_, toPrint := twitter.PublishTweet(tweet)

			c.Print(toPrint, "\n")

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

			tweets := twitter.GetTweets()

			c.Println(tweets[idAux].PrintableTweet())

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := twitter.GetTweets()

			for i := 0; i < len(tweets); i++ {
				c.Println(tweets[i].PrintableTweet())
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showUserTweets",
		Help: "Shows a users tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write the user: ")

			user := c.ReadLine()
			tweets := twitter.GetTweetsByUser(user)

			for i := 0; i < len(tweets); i++ {
				c.Println(tweets[i].PrintableTweet())
			}

			c.Println(tweets)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showUser",
		Help: "Shows a user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)
			c.Print("Write the user: ")

			user := c.ReadLine()

			userOut := twitter.GetUser(user)

			c.Println(userOut)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "registerUser",
		Help: "registers a user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write the user: ")

			user := c.ReadLine()
			c.Print("Write the mail: ")
			mail := c.ReadLine()
			c.Print("Write the nick: ")
			nick := c.ReadLine()
			c.Print("Write the password: ")
			password := c.ReadLine()

			userOut := domain.NewUser(user, mail, nick, password)

			twitter.Register(userOut)

			c.Println("succes")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "logIn",
		Help: "logs in a user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write the mail or nick: ")
			mail := c.ReadLine()

			c.Print("Write the password: ")
			password := c.ReadLine()

			twitter.LogIn(mail, password)

			c.Println("succes")

			return
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "logOut",
		Help: "logs out a user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write the user: ")
			user := c.ReadLine()

			twitter.LogOut(user)

			c.Println("succes")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "eraseTweet",
		Help: "deletes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write the user: ")
			user := c.ReadLine()

			c.Print("Write the tweet: ")
			tweet := c.ReadLine()

			twitter.Erase(user, tweet)

			c.Println("succes")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "modifyTweet",
		Help: "modifies a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write the user: ")
			user := c.ReadLine()

			c.Print("Write the new tweet: ")
			tweet := c.ReadLine()

			c.Print("Write the tweet id: ")

			id := c.ReadLine()
			idAux, _ := strconv.Atoi(id)

			twitter.Modify(idAux, tweet, user)

			c.Println("succes")

			return
		},
	})

	shell.Run()

}
