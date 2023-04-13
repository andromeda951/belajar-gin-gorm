package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/books/:id/:title", booksHandler) // localhost:8080/books/1/Golang
	router.GET("/query", queryHandler)            // localhost:8080/query?title=Belajar Gin&price=100000
	router.POST("/books", postBooksHandler)

	router.Run()
}

func rootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Andromeda",
		"bio":  "Software Engineer",
	})
}

func helloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"title":    "Hello World",
		"subtitle": "Belajar Gin dan Gorm",
	})
}

func booksHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	title := ctx.Param("title")

	ctx.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func queryHandler(ctx *gin.Context) {
	title := ctx.Query("title")
	price := ctx.Query("price")

	ctx.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

type BookInput struct {
	Title    string
	Price    int
	SubTitle string `json:"sub_title"`
}

func postBooksHandler(ctx *gin.Context) {
	var bookInput BookInput

	err := ctx.ShouldBindJSON(&bookInput)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"title":     bookInput.Title,
		"price":     bookInput.Price,
		"sub_title": bookInput.SubTitle,
	})
}
