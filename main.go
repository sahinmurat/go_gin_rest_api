package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articles = []Article{
	{ID: "1", Title: "First Article Title", Content: "First Content"},
	{ID: "2", Title: "Second Article Title", Content: "Second Content"},
}

func getAllArticles(c *gin.Context) {
	c.JSON(http.StatusOK, articles)
}

func main() {
	router := gin.Default()

	router.GET("/articles", getAllArticles)

	router.Run("localhost:8080")
}
