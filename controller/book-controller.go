package controller

import (
	"github.com/firattamur/go-web-with-gin/models"
	"github.com/firattamur/go-web-with-gin/service"
	"github.com/gin-gonic/gin"
)

type BookController interface {
	FindAll() []models.Book
	Save(ctx *gin.Context) models.Book
}

type controller struct {
	service service.BookService
}

func New(service service.BookService) BookController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []models.Book {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) models.Book {
	var book models.Book

	ctx.BindJSON(&book)
	c.service.Save(book)

	return book
}
