package main

import (
	"Tienda/internal/service"
	"Tienda/internal/store"
	"Tienda/internal/transport"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"fmt"
	"log"
	"net/http"
)

func main() {

	// Conectar a SQLite
	db, err := sql.Open("sqlite3", "./productos.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Crear el table si no existe
	q := `
CREATE TABLE IF NOT EXISTS productos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		category TEXT NOT NULL
	)
`

	if _, err := db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}

	// Inyectar nuestras dependencias
	productosStore := store.New(db)
	productoService := service.New(productosStore)
	productoHandler := transport.New(productoService)

	// Configurar las rutas
	http.HandleFunc("/productos", productoHandler.HandleProductos)
	http.HandleFunc("/productos/", productoHandler.HandleProductoByID)

	fmt.Println("🚀 Servidor ejecutandose en http://localhost:8000")
	fmt.Println("📚 API Endpoints:")
	fmt.Println("  - GET /productos			 - Obtener todos los productos")
	fmt.Println("  - POST /productos		 - Crear un nuevo producto")
	fmt.Println("  - GET /productos/{id}	 - Obtener un producto")
	fmt.Println("  - PUT /productos/{id}	 - Actualizar un producto")
	fmt.Println("  - DELETE /productos/{id}	 - Eliminar un producto")

	// Empezar y escuchar al servidor
	log.Fatal(http.ListenAndServe(":8000", nil))

}

/*
tienda/
	main.go
	internal/
		model/
			producto.go
		store/
			producto_store.go
		service/
			producto_service.go
		transport/
			producto_handler.go

model: Define nuestras estructuras de datos
store: maneja la base de datos
service: contiene la logica del negocio
transport: maneja http requests/responses

*/
