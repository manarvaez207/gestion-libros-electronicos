package main

import (
	"log"
	"net/http"

	"sistema-libros/internal/web"
)

func main() {
	app := web.NuevoApp("data/libros.json")

	log.Println("Sistema de Gestión de Libros Electrónicos - API Web")
	log.Println("Servidor iniciado en http://localhost:8080")

	if err := http.ListenAndServe(":8080", app.Rutas()); err != nil {
		log.Fatal(err)
	}
}
