package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var songs []SongResponse

func GetSongsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, songs)
}
