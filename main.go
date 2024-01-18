package main

import (
	"fmt"
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
	fmt.Println("all article")
	c.JSON(http.StatusOK, articles)
}

func getOneArticle(c *gin.Context) {
	fmt.Println("one article")
	id := c.Param("id")
	for _, article := range articles {
		if article.ID == id {
			c.JSON(http.StatusOK, article)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
}

func main() {
	router := gin.Default()

	router.GET("/articles", getAllArticles)
	router.GET("/articles/:id", getOneArticle)

	router.Run("localhost:8080")
}
