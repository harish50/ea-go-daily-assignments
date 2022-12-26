package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestUpdateSong(t *testing.T) {
	s := gin.New()
	s.PUT("/api/songs/:id", UpdateSong)
	releaseDate, _ := time.Parse(time.RFC3339, "2012-04-23T18:25:43.511Z")

	songs = append(songs, SongResponse{
		ID:           1,
		Title:        "Cheap thrills",
		Singer:       "Sia",
		Writer:       "Sia",
		ReleasedDate: releaseDate})

	reqBody := `{
		"title": "Unstoppable"
	}`
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/songs/1", strings.NewReader(reqBody))
	s.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	resBody, _ := io.ReadAll(res.Body)
	response := SongResponse{}
	json.Unmarshal(resBody, &response)
	assert.Equal(t, "Unstoppable", response.Title)
}
