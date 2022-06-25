package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/detail", detailHandler)
	router.GET("/book/:id", bookHandler)
	router.GET("/books", queryHandler)
	router.Run(":8080")
}

func rootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"pesan": "OK",
	})
}

func detailHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"pesan": "Detail",
	})
}

func bookHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func queryHandler(ctx *gin.Context) {
	title := ctx.Query("title")

	ctx.JSON(http.StatusOK, gin.H{"title": title})
}
