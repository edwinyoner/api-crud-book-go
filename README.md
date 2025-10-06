# 📚 API CRUD Book en Go

Este proyecto es una **API REST CRUD** desarrollada en **Go puro (sin frameworks)**, que permite gestionar un conjunto de libros en una base de datos **SQLite3**.  
El objetivo es aplicar las **buenas prácticas de arquitectura en Go**, utilizando una estructura modular y escalable.

---

## 🏗️ Arquitectura del proyecto

El proyecto sigue una estructura limpia basada en capas (`internal/`) y la convención estándar de proyectos Go:

```bash
api-crud-book-go/
│
├── cmd/
│   └── api/
│       └── main.go              # Punto de entrada principal
│
├── internal/
│   ├── app/
│   │   ├── server.go            # Inicialización del servidor HTTP
│   │   └── routes.go            # Definición de rutas
│   │
│   ├── config/
│   │   └── config.go            # Configuración y conexión a la base de datos
│   │
│   ├── model/
│   │   └── book.go              # Modelo Book
│   │
│   ├── service/
│   │   └── book_service.go      # Lógica de negocio
│   │
│   ├── store/
│   │   └── sqlite_store.go      # Acceso y operaciones a la base de datos SQLite
│   │
│   └── transport/
│       └── http/
│           ├── handler.go       # Handlers de las rutas
│           └── middleware.go    # Middlewares personalizados
│
├── go.mod
├── go.sum
├── books.db                     # Base de datos local (SQLite)
└── README.md
