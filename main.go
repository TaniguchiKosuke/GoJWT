package main

import (
	"GoJWT/auth"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Post struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

func index(c *gin.Context) {
	post := &Post{
		Title: "test title",
		Content: "test content",
	}
	c.JSON(200, *post)
}

func indexV2(w http.ResponseWriter, r *http.Request) {
	post := &Post{
		Title: "test title v2",
		Content: "test content v2",
    }
    json.NewEncoder(w).Encode(post)
}

func main() {
	router := gin.Default()
	router.GET("/index", index)
	router.GET("/token", auth.GetTokenHandler)
	checkToken := auth.CheckToken()
	router.GET("/private/index", gin.WrapH(checkToken.Handler(indexV2)))
	router.Run(":8000")
}