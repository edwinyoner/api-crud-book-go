package service

import (
	"github.com/edwinyoner/api-crud-book-go/internal/model"
	"github.com/edwinyoner/api-crud-book-go/internal/store"
)

type BookService struct {
	store *store.BookStore
}

func New(store *store.BookStore) *BookService {
	return &BookService{store: store}
}

func (s *BookService) GetAllBooks() ([]model.Book, error) {
	return s.store.GetAll()
}

func (s *BookService) AddBook(book model.Book) error {
	return s.store.Create(book)
}
