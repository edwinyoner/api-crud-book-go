package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./../../books.db")
	if err != nil {
		log.Fatalf("Error al abrir la base de datos: %v", err)
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS libros (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		titulo TEXT NOT NULL,
		autor TEXT NOT NULL
	);
	`
	if _, err := db.Exec(createTableQuery); err != nil {
		log.Fatalf("Error al crear la tabla: %v", err)
	}

	return db
}
