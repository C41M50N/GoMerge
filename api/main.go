package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CounterResult struct {
	Count uint16 `json:"count"`
}

var count uint16 = 0

func incrementCounter() { count += 1 }

func main() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowMethods: []string{"GET", "POST"},
	}))

	router.GET("/counter", func(ctx *gin.Context) {
		defer incrementCounter()
		fmt.Printf("Count: %d\n", count)
		ctx.JSON(200, CounterResult{Count: count})
	})

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "success",
		})
	})

	fmt.Println("Starting Server ...")
	router.Run()
}
