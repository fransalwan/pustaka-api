package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello/:name", helloHandler)
	router.GET("/books/:id", booksHandler)

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
