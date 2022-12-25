package day9_songs

import "github.com/gin-gonic/gin"

func main() {
	s := gin.New()
	s.POST("/api/songs", CreateHandler)
}
