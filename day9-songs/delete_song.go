package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteSong(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	d := -1
	for i, s := range songs {
		if s.ID == id {
			d = i
			break
		}
	}
	if d == -1 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Song not found"})
		return
	}
	songs = append(songs[:d], songs[d+1:]...)
}
