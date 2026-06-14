package reportes

import (
	"errors"
	"fmt"
	"strings"
)

// LibroReporte define los métodos necesarios para generar reportes.
// Se utiliza una interfaz para aplicar polimorfismo y evitar depender directamente del struct Libro.
type LibroReporte interface {
	ID() int
	Titulo() string
	Autor() string
	Categoria() string
	Anio() int
	Disponible() bool
	Descripcion() string
}

// ContarDisponibles cuenta cuántos libros están disponibles.
func ContarDisponibles(lista []LibroReporte) int {
	total := 0

	for _, libro := range lista {
		if libro.Disponible() {
			total++
		}
	}

	return total
}

// ContarNoDisponibles cuenta cuántos libros no están disponibles.
func ContarNoDisponibles(lista []LibroReporte) int {
	total := 0

	for _, libro := range lista {
		if !libro.Disponible() {
			total++
		}
	}

	return total
}

// ContarPorCategoria cuenta los libros que pertenecen a una categoría específica.
func ContarPorCategoria(lista []LibroReporte, categoria string) (int, error) {
	categoria = strings.TrimSpace(strings.ToLower(categoria))

	if categoria == "" {
		return 0, errors.New("la categoría para el reporte no puede estar vacía")
	}

	total := 0

	for _, libro := range lista {
		if strings.ToLower(libro.Categoria()) == categoria {
			total++
		}
	}

	return total, nil
}

// GenerarResumen crea un resumen general del catálogo de libros.
func GenerarResumen(lista []LibroReporte) string {
	totalLibros := len(lista)
	disponibles := ContarDisponibles(lista)
	noDisponibles := ContarNoDisponibles(lista)

	return fmt.Sprintf(
		"Resumen del sistema:\nTotal de libros: %d\nLibros disponibles: %d\nLibros no disponibles: %d",
		totalLibros,
		disponibles,
		noDisponibles,
	)
}
