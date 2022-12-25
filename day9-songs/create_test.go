package day9_songs

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

func makeACall(reqBody string) *httptest.ResponseRecorder {
	s := gin.New()
	s.POST("/api/songs", CreateHandler)

	req, _ := http.NewRequest("POST", "/api/songs", strings.NewReader(reqBody))
	w := httptest.NewRecorder()
	s.ServeHTTP(w, req)
	return w
}

func TestCreateHandler(t *testing.T) {
	reqBody := `{
		"title" : "Cheap thrills",
		"singer" : "Sia",
		"writer" : "Sia",
		"released_date": "2012-04-23T18:25:43.511Z"
	}`
	res := makeACall(reqBody)

	assert.Equal(t, http.StatusCreated, res.Code)
	resBody, _ := io.ReadAll(res.Body)
	response := SongResponse{}
	json.Unmarshal(resBody, &response)
	assert.Equal(t, 1, response.ID)
	tm, _ := time.Parse(time.RFC3339, "2012-04-23T18:25:43.511Z")
	assert.Equal(t, tm, response.ReleasedDate)
}
