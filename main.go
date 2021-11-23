package main

import (
	"github.com/firattamur/go-web-with-gin/controller"
	"github.com/firattamur/go-web-with-gin/service"
	"github.com/gin-gonic/gin"
)

var (
	bookService    service.BookService       = service.New()
	bookController controller.BookController = controller.New(bookService)
)

func main() {

	app := gin.Default()

	app.GET("/books", func(ctx *gin.Context) {
		ctx.JSON(200, bookController.FindAll())
	})

	app.POST("/books", func(ctx *gin.Context) {
		ctx.JSON(200, bookController.Save(ctx))
	})

	app.Run(":8000")

}
