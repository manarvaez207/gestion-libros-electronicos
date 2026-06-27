package busquedas

import "testing"

type libroPrueba struct {
	titulo    string
	autor     string
	categoria string
}

func (l libroPrueba) Titulo() string {
	return l.titulo
}

func (l libroPrueba) Autor() string {
	return l.autor
}

func (l libroPrueba) Categoria() string {
	return l.categoria
}

func (l libroPrueba) Descripcion() string {
	return l.titulo + " - " + l.autor
}

func TestBuscarPorTitulo(t *testing.T) {
	catalogo := []LibroBuscable{
		libroPrueba{
			titulo:    "Introduccion a Go",
			autor:     "Alan Donovan",
			categoria: "Programacion",
		},
		libroPrueba{
			titulo:    "Clean Code",
			autor:     "Robert C. Martin",
			categoria: "Buenas practicas",
		},
	}

	resultados, err := BuscarPorTitulo(catalogo, "go")

	if err != nil {
		t.Fatalf("no se esperaba error, pero se obtuvo: %v", err)
	}

	if len(resultados) != 1 {
		t.Fatalf("se esperaba 1 resultado, se obtuvieron %d", len(resultados))
	}
}

func TestBuscarPorCategoria(t *testing.T) {
	catalogo := []LibroBuscable{
		libroPrueba{
			titulo:    "Introduccion a Go",
			autor:     "Alan Donovan",
			categoria: "Programacion",
		},
		libroPrueba{
			titulo:    "Clean Code",
			autor:     "Robert C. Martin",
			categoria: "Buenas practicas",
		},
	}

	resultados, err := BuscarPorCategoria(catalogo, "programacion")

	if err != nil {
		t.Fatalf("no se esperaba error, pero se obtuvo: %v", err)
	}

	if len(resultados) != 1 {
		t.Fatalf("se esperaba 1 resultado, se obtuvieron %d", len(resultados))
	}
}

func TestBuscarPorTituloVacio(t *testing.T) {
	catalogo := []LibroBuscable{}

	_, err := BuscarPorTitulo(catalogo, "")

	if err == nil {
		t.Fatal("se esperaba error cuando el texto de busqueda esta vacio")
	}
}
