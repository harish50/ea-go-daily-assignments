package main

import "github.com/gin-gonic/gin"

func main() {
	s := gin.New()
	s.POST("/api/songs", CreateSongHandler)
	s.GET("/api/songs", GetSongsHandler)
	s.Run("localhost:8080")
}
