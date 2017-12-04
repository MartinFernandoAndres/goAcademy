package rest

import (
	"github.com/goAcademy/src/service"
	"github.com/goAcademy/src/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)


type GinTweet struct {
	User domain.User
	Text string
	URL  string
	ID   int
}
type GinServer struct {
	tweetManager *service.Twitter
}

func NewGinServer (tweetManager *service.Twitter) *GinServer{
	return &GinServer{tweetManager}
}

func (server *GinServer) StartGinServer() {

	router := gin.Default()

	router.GET("/listTweets", server.getTweets)
	router.GET("/listTweets/:user", server.getTweetsByUser)
	router.POST("publishTextTweet", server.publishTweet)
	router.POST("publishImageTweet", server.publishImageTweet)
	router.POST("publishTweet", server.publishImageTweet)

	go router.Run()
}

func (server *GinServer) getTweets(c *gin.Context) {
	c.JSON(http.StatusOK, server.tweetManager.GetTweets())
}


func (server *GinServer) getTweetsByUser(c *gin.Context) {

	user := c.Param("user")
	c.JSON(http.StatusOK, server.tweetManager.GetTweetsByUser(user))
}

func (server *GinServer) publishTweet(c *gin.Context) {

	//quit := make(chan bool)

	var tweetdata GinTweet
	c.Bind(&tweetdata)

	tweetToPublish := domain.NewTextTweet(tweetdata.User, tweetdata.Text)

	id, err := server.tweetManager.PublishTweet(tweetToPublish)//hay que hacer chanels

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error publishing tweet "+err.Error())
	} else {
		c.JSON(http.StatusOK, struct{ Id int }{id})
	}
}

func (server *GinServer) publishImageTweet(c *gin.Context) {

	//quit := make(chan bool)

	var tweetdata GinTweet
	c.Bind(&tweetdata)

	tweetToPublish := domain.NewImageTweet(tweetdata.User, tweetdata.Text, tweetdata.URL)

	id, err := server.tweetManager.PublishTweet(tweetToPublish)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error publishing tweet "+err.Error())
	} else {
		c.JSON(http.StatusOK, struct{ Id int }{id})
	}
}

func (server *GinServer) publishQuoteTweet(c *gin.Context) {

	//quit := make(chan bool)

	var tweetdata GinTweet
	c.Bind(&tweetdata)

	quotedTweet := server.tweetManager.GetTweetById(tweetdata.ID)
	tweetToPublish := domain.NewQuoteTweet(tweetdata.User, tweetdata.Text, quotedTweet)

	id, err := server.tweetManager.PublishTweet(tweetToPublish)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error publishing tweet "+err.Error())
	} else {
		c.JSON(http.StatusOK, struct{ Id int }{id})
	}
}
