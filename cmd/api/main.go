package main

import (
	"log"

	"github.com/edwinyoner/api-crud-book-go/internal/app"
)

func main() {
	server := app.NewServer()
	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
