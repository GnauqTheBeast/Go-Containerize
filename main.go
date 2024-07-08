package main

import (
	"context"
	"fmt"
	"test/ent"
	"test/ent/user"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	client *ent.Client
	err    error
)

func main() {
	client, err = ent.Open("postgres", "host=postgres port=5433 user=postgres dbname=test password=postgres sslmode=disable")
	if err != nil {
		fmt.Printf("failed opening connection to postgres: %v", err)
		return
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		fmt.Printf("failed creating schema resources: %v", err)
	}

	r := gin.Default()
	r.GET("/ping", PingHandler)
	r.Run() // listen and serve on 0.0.0.0:8080

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run()

}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name("quang")).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	return u, nil
}

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"users": "test nhe ae",
	})
}
