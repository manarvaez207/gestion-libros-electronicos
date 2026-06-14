package busquedas

import (
	"errors"
	"strings"
)

// LibroBuscable define los métodos necesarios para buscar libros.
// Se utiliza una interfaz para permitir polimorfismo y desacoplar el módulo de búsquedas.
type LibroBuscable interface {
	Titulo() string
	Autor() string
	Categoria() string
	Descripcion() string
}

// BuscarPorTitulo busca libros por coincidencia parcial en el título.
func BuscarPorTitulo(lista []LibroBuscable, texto string) ([]LibroBuscable, error) {
	texto = strings.TrimSpace(strings.ToLower(texto))

	if texto == "" {
		return nil, errors.New("el texto de búsqueda no puede estar vacío")
	}

	resultados := make([]LibroBuscable, 0)

	for _, libro := range lista {
		titulo := strings.ToLower(libro.Titulo())

		if strings.Contains(titulo, texto) {
			resultados = append(resultados, libro)
		}
	}

	if len(resultados) == 0 {
		return nil, errors.New("no se encontraron libros con ese título")
	}

	return resultados, nil
}

// BuscarPorCategoria busca libros por coincidencia parcial en la categoría.
func BuscarPorCategoria(lista []LibroBuscable, texto string) ([]LibroBuscable, error) {
	texto = strings.TrimSpace(strings.ToLower(texto))

	if texto == "" {
		return nil, errors.New("la categoría de búsqueda no puede estar vacía")
	}

	resultados := make([]LibroBuscable, 0)

	for _, libro := range lista {
		categoria := strings.ToLower(libro.Categoria())

		if strings.Contains(categoria, texto) {
			resultados = append(resultados, libro)
		}
	}

	if len(resultados) == 0 {
		return nil, errors.New("no se encontraron libros en esa categoría")
	}

	return resultados, nil
}
