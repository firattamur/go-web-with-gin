package service

import "github.com/firattamur/go-web-with-gin/models"

type BookService interface {
	Save(models.Book) models.Book
	FindAll() []models.Book
}

type bookService struct {
	books []models.Book
}

func New() BookService {
	return &bookService{}
}

func (service *bookService) Save(book models.Book) models.Book {
	service.books = append(service.books, book)
	return book
}

func (service *bookService) FindAll() []models.Book {
	return service.books
}
