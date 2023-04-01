package main

import(
	"bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
	"github.com/rs/xid"
    "github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine{
    router := gin.Default()
    return router
}

func TestGetAlbum(t *testing.T){
	r := SetUpRouter()
	r.GET("/albums", getAlbums)
	req, _ := http.NewRequest("GET", "/albums", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var albums []album
	json.Unmarshal(w.Body.Bytes(), &albums)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, albums)
}

func TestPostAlbums(t *testing.T){
	r := SetUpRouter()
	r.POST("/albums", postAlbums)
	albumId := xid.New().String()

	albums := album {
		ID : albumId,
		Title : "Test Album",
		Artist : "Test Artist",
		Price : 29.99,
	}
	jsonValue, _ := json.Marshal(albums)
	req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}