package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"testing"
)

func Test_ResolveConfig(t *testing.T) {
	rawConfig := getRawConfig()
	resolvedConfig, err := resolveRawConfig(rawConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Configuration errors - %s\n", err)
		os.Exit(1)
	}

	var expected = &config{}

	if exp, got := expected, resolvedConfig; exp != got {
		t.Fatalf("unexpected results for query, expected %s, got %s", exp, got)
	}
}

func Test_SetUpServer(t *testing.T) {
	// Read in a raft node configuration
	rawConfig := getRawConfig()
	config, err := resolveRawConfig(rawConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Configuration errors - %s\n", err)
		os.Exit(1)
	}

	serverAddr = config.RaftAddress.String()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Configuration errors - %s\n", err)
		os.Exit(1)
	}

	router := gin.Default()
	router.Use(CORSMiddleware())
	rg := router.Group("api/v1/")
	{
		rg.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "Pong")
		})
		rg.POST("/register", postNewUser)
		rg.POST("/login", postNewLogin)
		rg.POST("/createTweet", postNewTweet)
		rg.POST("/followUser", postNewFollow)
		rg.POST("/unfollowUser", postNewUnfollow)
		rg.GET("/getTweetsByUser/:id", getTweetsByUser)
		rg.GET("/getFeedByUser/:id", getFeedByUser)
		rg.GET("/getFollowedByUser/:id", getFollowedByUser)
		rg.GET("/getFollowingByUser/:id", getFollowingByUser)
		rg.GET("/getUser/:id", getUser)
		rg.GET("/getUsers", getUsers)
		rg.GET("/getUsersNotFollowed/:id", getUsersNotFollowed)
	}

	if err := router.Run(config.ServerAddress.String()); err != nil {
		log.Fatalf("could not run server: %v", err)
		return
	}
}
