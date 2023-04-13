package main

import (
	"belajar-gin-gorm/handler"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	}

	fmt.Println("Database connection succeed")

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler) // localhost:8080/books/1/Golang
	v1.GET("/query", handler.QueryHandler)            // localhost:8080/query?title=Belajar Gin&price=100000
	v1.POST("/books", handler.PostBooksHandler)

	router.Run()
}
