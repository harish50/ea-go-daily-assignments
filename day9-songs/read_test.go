package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestReadHandler(t *testing.T) {
	s := gin.New()
	s.GET("/api/songs", ReadHandler)
	releaseDate, _ := time.Parse(time.RFC3339, "2012-04-23T18:25:43.511Z")

	songs = append(songs, SongResponse{
		ID:           1,
		Title:        "Cheap thrills",
		Singer:       "Sia",
		Writer:       "Sia",
		ReleasedDate: releaseDate})

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/songs", nil)
	s.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	resBody, _ := io.ReadAll(res.Body)
	var response []SongResponse
	json.Unmarshal(resBody, &response)
	assert.Equal(t, 1, response[0].ID)
	tm, _ := time.Parse(time.RFC3339, "2012-04-23T18:25:43.511Z")
	assert.Equal(t, tm, response[0].ReleasedDate)
}
