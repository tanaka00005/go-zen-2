package main

import (
	"fmt"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/tanaka00005/api/internal/presenter"
)

func main() {
	srv := presenter.NewServer()

	if err := srv.Run(context.Background()); err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/",func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message":"hello world",
		})
	})
    fmt.Println("Hello World !")

	r.Run(":8080")
}