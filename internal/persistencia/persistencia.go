package persistencia

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"sistema-libros/internal/libros"
)

// libroJSON representa la estructura que se guardará en el archivo JSON.
// Se usa este struct auxiliar porque Libro tiene atributos privados por encapsulación.
type libroJSON struct {
	ID         int    `json:"id"`
	Titulo     string `json:"titulo"`
	Autor      string `json:"autor"`
	Categoria  string `json:"categoria"`
	Anio       int    `json:"anio"`
	Disponible bool   `json:"disponible"`
}

// GuardarLibros guarda una lista de libros en un archivo JSON.
func GuardarLibros(ruta string, lista []libros.Libro) error {
	if strings.TrimSpace(ruta) == "" {
		return errors.New("la ruta del archivo no puede estar vacía")
	}

	registros := make([]libroJSON, 0, len(lista))

	for _, libro := range lista {
		registros = append(registros, libroJSON{
			ID:         libro.ID(),
			Titulo:     libro.Titulo(),
			Autor:      libro.Autor(),
			Categoria:  libro.Categoria(),
			Anio:       libro.Anio(),
			Disponible: libro.Disponible(),
		})
	}

	datos, err := json.MarshalIndent(registros, "", "  ")
	if err != nil {
		return fmt.Errorf("error al convertir los libros a JSON: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(ruta), 0755); err != nil {
		return fmt.Errorf("error al crear el directorio de datos: %w", err)
	}

	if err := os.WriteFile(ruta, datos, 0644); err != nil {
		return fmt.Errorf("error al guardar el archivo JSON: %w", err)
	}

	return nil
}

// CargarLibros lee una lista de libros desde un archivo JSON.
func CargarLibros(ruta string) ([]libros.Libro, error) {
	if strings.TrimSpace(ruta) == "" {
		return nil, errors.New("la ruta del archivo no puede estar vacía")
	}

	datos, err := os.ReadFile(ruta)
	if errors.Is(err, os.ErrNotExist) {
		return []libros.Libro{}, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error al leer el archivo JSON: %w", err)
	}

	if strings.TrimSpace(string(datos)) == "" {
		return []libros.Libro{}, nil
	}

	var registros []libroJSON

	if err := json.Unmarshal(datos, &registros); err != nil {
		return nil, fmt.Errorf("error al convertir el JSON a libros: %w", err)
	}

	lista := make([]libros.Libro, 0, len(registros))

	for _, registro := range registros {
		libro, err := libros.NuevoLibro(
			registro.ID,
			registro.Titulo,
			registro.Autor,
			registro.Categoria,
			registro.Anio,
		)

		if err != nil {
			return nil, fmt.Errorf("libro inválido con ID %d: %w", registro.ID, err)
		}

		libro.CambiarDisponibilidad(registro.Disponible)
		lista = append(lista, libro)
	}

	return lista, nil
}
