package libros

import (
	"errors"
	"fmt"
)

// Libro representa un libro electrónico dentro del sistema.
// Sus atributos están en minúscula para aplicar encapsulación,
// evitando que se modifiquen directamente desde otros paquetes.
type Libro struct {
	id         int
	titulo     string
	autor      string
	categoria  string
	anio       int
	disponible bool
}

// NuevoLibro funciona como constructor.
// Permite crear un libro validando primero los datos principales.
func NuevoLibro(id int, titulo string, autor string, categoria string, anio int) (Libro, error) {
	libro := Libro{
		id:         id,
		disponible: true,
	}

	if err := libro.SetTitulo(titulo); err != nil {
		return Libro{}, err
	}

	if err := libro.SetAutor(autor); err != nil {
		return Libro{}, err
	}

	if err := libro.SetCategoria(categoria); err != nil {
		return Libro{}, err
	}

	if err := libro.SetAnio(anio); err != nil {
		return Libro{}, err
	}

	return libro, nil
}

// Métodos getter.
// Permiten consultar los valores privados del struct.
func (l Libro) ID() int {
	return l.id
}

func (l Libro) Titulo() string {
	return l.titulo
}

func (l Libro) Autor() string {
	return l.autor
}

func (l Libro) Categoria() string {
	return l.categoria
}

func (l Libro) Anio() int {
	return l.anio
}

func (l Libro) Disponible() bool {
	return l.disponible
}

// Métodos setter.
// Usan receptor puntero para modificar directamente el libro original.
func (l *Libro) SetTitulo(titulo string) error {
	if titulo == "" {
		return errors.New("el título del libro no puede estar vacío")
	}

	l.titulo = titulo
	return nil
}

func (l *Libro) SetAutor(autor string) error {
	if autor == "" {
		return errors.New("el autor del libro no puede estar vacío")
	}

	l.autor = autor
	return nil
}

func (l *Libro) SetCategoria(categoria string) error {
	if categoria == "" {
		return errors.New("la categoría del libro no puede estar vacía")
	}

	l.categoria = categoria
	return nil
}

func (l *Libro) SetAnio(anio int) error {
	if anio <= 0 {
		return errors.New("el año del libro debe ser mayor a cero")
	}

	if anio > 2026 {
		return errors.New("el año del libro no puede ser mayor al año actual")
	}

	l.anio = anio
	return nil
}

// CambiarDisponibilidad permite marcar si el libro está disponible o no.
func (l *Libro) CambiarDisponibilidad(disponible bool) {
	l.disponible = disponible
}

// Descripcion retorna una cadena con la información principal del libro.
func (l Libro) Descripcion() string {
	estado := "No disponible"

	if l.disponible {
		estado = "Disponible"
	}

	return fmt.Sprintf(
		"ID: %d | Título: %s | Autor: %s | Categoría: %s | Año: %d | Estado: %s",
		l.id,
		l.titulo,
		l.autor,
		l.categoria,
		l.anio,
		estado,
	)
}
