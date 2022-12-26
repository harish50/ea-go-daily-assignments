package main

import "github.com/gin-gonic/gin"

func main() {
	s := gin.New()
	s.POST("/api/songs", CreateHandler)
	s.GET("/api/songs", ReadHandler)
	s.Run("localhost:8080")
}
