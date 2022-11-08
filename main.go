package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "Harry Potter", Author: "J.K Rowling"},
	{ID: "2", Title: "The Lord Of The Rings", Author: "J.R.R Tolkein"},
	{ID: "3", Title: "The Wizard Of Oz", Author: "L. Frank Baum"},
}

func listBooksHandler(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func createBookHandler(c *gin.Context) {
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	books = append(books, book)
	c.JSON(http.StatusCreated, book)
}

func deleteBookHandler(c *gin.Context) {
	id := c.Param("id")

	for i, a := range books {
		if a.ID == id {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}

	c.Status(http.StatusNoContent)
}

func main() {
	r := gin.New()

	// GET Request
	r.GET("/books", listBooksHandler)

	// POST Request
	r.POST("/books", createBookHandler)

	// DELETE Request
	r.DELETE("/books/:id", deleteBookHandler)

	r.Run()
}
