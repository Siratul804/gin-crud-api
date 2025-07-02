package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     int    ⁠ json:"id" ⁠
	Title  string ⁠ json:"title" ⁠
	Author string ⁠ json:"author" ⁠
	Year   int    ⁠ json:"year" ⁠
}

var books = []Book{}
var nextID = 1

func main() {
	r := gin.Default()

	r.GET("/books", getBooks)
	r.GET("/books/:id", getBook)
	r.POST("/books", createBook)
	r.PUT("/books/:id", updateBook)
	r.DELETE("/books/:id", deleteBook)

	r.Run(":8080")
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func getBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, b := range books {
		if b.ID == id {
			c.JSON(http.StatusOK, b)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
}

func createBook(c *gin.Context) {
	var b Book
	if c.BindJSON(&b) == nil {
		b.ID = nextID
		nextID++
		books = append(books, b)
		c.JSON(http.StatusCreated, b)
	}
}

func updateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var u Book
	if c.BindJSON(&u) == nil {
		for i, b := range books {
			if b.ID == id {
				u.ID = id
				books[i] = u
				c.JSON(http.StatusOK, u)
				return
			}
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
}

func deleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
}