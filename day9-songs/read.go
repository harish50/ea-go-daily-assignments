package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var songs []SongResponse

func ReadHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, songs)
}
