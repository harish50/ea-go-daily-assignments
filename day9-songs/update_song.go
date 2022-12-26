package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateSong(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	request := SongRequest{}
	err1 := ctx.BindJSON(&request)
	if err1 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	for i, s := range songs {
		if s.ID == id {
			if len(request.Title) != 0 {
				songs[i].Title = request.Title
			}
			if len(request.Singer) != 0 {
				songs[i].Singer = request.Singer
			}
			if len(request.Writer) != 0 {
				songs[i].Writer = request.Writer
			}
			if !request.ReleasedDate.IsZero() {
				songs[i].ReleasedDate = request.ReleasedDate
			}
			ctx.JSON(http.StatusOK, songs[i])
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Song not found"})
}
