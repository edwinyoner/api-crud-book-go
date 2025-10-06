package app

import (
	"net/http"

	httpTransport "github.com/edwinyoner/api-crud-book-go/internal/transport/http"
)

func SetupRoutes(handler *httpTransport.BookHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/books", handler.HandleBooks)
	mux.HandleFunc("/book/", handler.HandleBookByID)
	return mux
}
