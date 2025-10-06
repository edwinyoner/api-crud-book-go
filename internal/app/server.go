package app

import (
	"log"
	"net/http"

	"github.com/edwinyoner/api-crud-book-go/internal/config"
	"github.com/edwinyoner/api-crud-book-go/internal/service"
	"github.com/edwinyoner/api-crud-book-go/internal/store"
	httpTransport "github.com/edwinyoner/api-crud-book-go/internal/transport/http"
)

type Server struct {
	router *http.ServeMux
}

func NewServer() *Server {
	db := config.InitDB()
	bookStore := store.New(db)
	bookService := service.New(bookStore)
	bookHandler := httpTransport.NewBookHandler(bookService)

	router := SetupRoutes(bookHandler)

	return &Server{router: router}
}

func (s *Server) Run(addr string) error {
	log.Printf("Servidor escuchando en %s", addr)
	return http.ListenAndServe(addr, s.router)
}
