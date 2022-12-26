package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type SongRequest struct {
	Title        string    `json:"title"`
	Singer       string    `json:"singer"`
	Writer       string    `json:"writer"`
	ReleasedDate time.Time `json:"released_date"`
}

type SongResponse struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Singer       string    `json:"singer"`
	Writer       string    `json:"writer"`
	ReleasedDate time.Time `json:"released_date"`
}

func CreateHandler(ctx *gin.Context) {
	request := SongRequest{}
	err := ctx.BindJSON(&request)
	if err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, nil)
	}

	response := SongResponse{
		ID:           len(songs) + 1,
		Title:        request.Title,
		Singer:       request.Singer,
		Writer:       request.Writer,
		ReleasedDate: request.ReleasedDate,
	}

	songs = append(songs, response)

	ctx.JSON(http.StatusCreated, response)
}
