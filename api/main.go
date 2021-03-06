package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	pb "simple-twitter.com/backend/rpc/proto"
)

var serverAddr string = ""

func main() {
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func postNewUser(c *gin.Context) {
	var opts []grpc.DialOption
	var conn *grpc.ClientConn
	var err error
	opts = append(opts, grpc.WithInsecure())
	conn, err = grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewTwitterClient(conn)
	var newUser pb.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, e := client.CreateUser(context.Background(), &newUser)

	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Fatalf("Failed to create new user: %v", e)
	}

	c.JSON(http.StatusOK, res)
	return
}

func postNewLogin(c *gin.Context) {
	var opts []grpc.DialOption
	var conn *grpc.ClientConn
	var err error
	opts = append(opts, grpc.WithInsecure())
	conn, err = grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewTwitterClient(conn)
	var newUser pb.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, e := client.LoginUser(context.Background(), &newUser)

	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Fatalf("Failed to create new user: %v", e)
	}

	c.JSON(http.StatusOK, res)
	return
}

func postNewTweet(c *gin.Context) {
	var opts []grpc.DialOption
	var conn *grpc.ClientConn
	var err error
	opts = append(opts, grpc.WithInsecure())
	conn, err = grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewTwitterClient(conn)
	var newTweet pb.Tweet
	if err := c.BindJSON(&newTweet); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, e := client.CreateTweet(context.Background(), &newTweet)

	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Fatalf("Failed to create new tweet: %v", e)
	}

	c.JSON(http.StatusOK, res)
	return
}

func postNewFollow(c *gin.Context) {
	var opts []grpc.DialOption
	var conn *grpc.ClientConn
	var err error
	opts = append(opts, grpc.WithInsecure())
	conn, err = grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewTwitterClient(conn)
	var newFollow pb.Follow
	if err := c.BindJSON(&newFollow); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, e := client.FollowUser(context.Background(), &newFollow)

	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Fatalf("Failed to create new tweet: %v", e)
	}

	c.JSON(http.StatusOK, res)
	return
}

func postNewUnfollow(c *gin.Context) {
	var opts []grpc.DialOption
	var conn *grpc.ClientConn
	var err error
	opts = append(opts, grpc.WithInsecure())
	conn, err = grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewTwitterClient(conn)
	var newFollow pb.Follow
	if err := c.BindJSON(&newFollow); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, e := client.UnfollowUser(context.Background(), &newFollow)

	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Fatalf("Failed to unfollow user: %v", e)
	}

	c.JSON(http.StatusOK, res)
	return
}

func getTweetsByUser(c *gin.Context) {
	var opts []grpc.DialOption
	var conn *grpc.ClientConn
	var err error
	opts = append(opts, grpc.WithInsecure())
	conn, err = grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewTwitterClient(conn)
	res, e := client.GetTweetsByUser(context.Background(), &pb.User{
		Id: c.Param("id"),
	})

	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Fatalf("Failed to get user tweets: %v", e)
	}

	c.JSON(http.StatusOK, res)
	return
}

func getFeedByUser(c *gin.Context) {
	var opts []grpc.DialOption
	var conn *grpc.ClientConn
	var err error
	opts = append(opts, grpc.WithInsecure())
	conn, err = grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewTwitterClient(conn)
	res, e := client.GetFeedByUser(context.Background(), &pb.User{
		Id: c.Param("id"),
	})

	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Fatalf("Failed to get feed: %v", e)
	}

	c.JSON(http.StatusOK, res)
	return
}

func getFollowedByUser(c *gin.Context) {
	var opts []grpc.DialOption
	var conn *grpc.ClientConn
	var err error
	opts = append(opts, grpc.WithInsecure())
	conn, err = grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewTwitterClient(conn)
	res, e := client.GetFollowedByUser(context.Background(), &pb.User{
		Id: c.Param("id"),
	})

	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Fatalf("Failed to get followers: %v", e)
	}

	c.JSON(http.StatusOK, res)
	return
}

func getFollowingByUser(c *gin.Context) {
	var opts []grpc.DialOption
	var conn *grpc.ClientConn
	var err error
	opts = append(opts, grpc.WithInsecure())
	conn, err = grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewTwitterClient(conn)
	res, e := client.GetFollowingByUser(context.Background(), &pb.User{
		Id: c.Param("id"),
	})

	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Fatalf("Failed to get following: %v", e)
	}

	c.JSON(http.StatusOK, res)
	return
}

func getUser(c *gin.Context) {
	var opts []grpc.DialOption
	var conn *grpc.ClientConn
	var err error
	opts = append(opts, grpc.WithInsecure())
	conn, err = grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewTwitterClient(conn)
	res, e := client.GetUser(context.Background(), &pb.User{
		Id: c.Param("id"),
	})

	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Fatalf("Failed to get following: %v", e)
	}

	c.JSON(http.StatusOK, res)
	return
}

func getUsers(c *gin.Context) {
	var opts []grpc.DialOption
	var conn *grpc.ClientConn
	var err error
	opts = append(opts, grpc.WithInsecure())
	conn, err = grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewTwitterClient(conn)
	res, e := client.GetUsers(context.Background(), &pb.User{})

	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Fatalf("Failed to get following: %v", e)
	}

	c.JSON(http.StatusOK, res)
	return
}

func getUsersNotFollowed(c *gin.Context) {
	var opts []grpc.DialOption
	var conn *grpc.ClientConn
	var err error
	opts = append(opts, grpc.WithInsecure())
	conn, err = grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewTwitterClient(conn)
	res, e := client.GetUsersNotFollowed(context.Background(), &pb.User{
		Id: c.Param("id"),
	})

	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Fatalf("Failed to get following: %v", e)
	}

	c.JSON(http.StatusOK, res)
	return
}
