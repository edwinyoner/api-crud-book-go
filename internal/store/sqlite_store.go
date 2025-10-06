package store

import (
	"database/sql"

	"github.com/edwinyoner/api-crud-book-go/internal/model"
)

type BookStore struct {
	db *sql.DB
}

func New(db *sql.DB) *BookStore {
	return &BookStore{db: db}
}

func (s *BookStore) GetAll() ([]model.Book, error) {
	rows, err := s.db.Query("SELECT id, titulo, autor FROM libros")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var b model.Book
		if err := rows.Scan(&b.ID, &b.Titulo, &b.Autor); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func (s *BookStore) Create(book model.Book) error {
	_, err := s.db.Exec("INSERT INTO libros (titulo, autor) VALUES (?, ?)", book.Titulo, book.Autor)
	return err
}
