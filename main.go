package main

import (
	"belajar-gin-gorm/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler) // localhost:8080/books/1/Golang
	v1.GET("/query", handler.QueryHandler)            // localhost:8080/query?title=Belajar Gin&price=100000
	v1.POST("/books", handler.PostBooksHandler)

	router.Run()
}
