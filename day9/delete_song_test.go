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

func TestDeleteSong(t *testing.T) {
	s := gin.New()
	s.DELETE("/api/songs/:id", DeleteSong)
	s.GET("/api/songs", GetSongsHandler)

	songs = append(songs, SongResponse{
		ID:           1,
		Title:        "Cheap thrills",
		Singer:       "Sia",
		Writer:       "Sia",
		ReleasedDate: time.Now()})

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/songs/1", nil)
	s.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	res = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/songs", nil)
	s.ServeHTTP(res, req)
	resBody, _ := io.ReadAll(res.Body)
	var response []SongResponse
	json.Unmarshal(resBody, &response)
	assert.Equal(t, 0, len(response))
}
