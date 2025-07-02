package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello/:name", helloHandler)
	router.GET("/books/:id", booksHandler)
	router.GET("/query", queryHandler)
	router.POST("/books", postBooksHandler)

	router.Run()

}

func rootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func helloHandler(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"message": "Hello, " + name + "!",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "Book ID: " + id,
	})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")
	c.JSON(200, gin.H{
		"message": "Title: " + title + ", Price: " + price,
	})
}

type BookInput struct {
	Title    string
	Price    int
	SubTitle string `json:"sub_title"`
}

func postBooksHandler(c *gin.Context) {
	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"title":     bookInput.Title,
		"price":     bookInput.Price,
		"sub_title": bookInput.SubTitle,
	})
}
