package http

import (
	"encoding/json"
	"net/http"

	"github.com/edwinyoner/api-crud-book-go/internal/model"
	"github.com/edwinyoner/api-crud-book-go/internal/service"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) HandleBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		books, err := h.service.GetAllBooks()
		if err != nil {
			http.Error(w, "Error al obtener los libros", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(books)
	case http.MethodPost:
		var book model.Book
		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}
		if err := h.service.AddBook(book); err != nil {
			http.Error(w, "Error al guardar el libro", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func (h *BookHandler) HandleBookByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Funcionalidad de buscar por ID en construcción"))
}
