package libros

import "testing"

func TestNuevoLibroValido(t *testing.T) {
	libro, err := NuevoLibro(1, "Introduccion a Go", "Alan Donovan", "Programacion", 2016)

	if err != nil {
		t.Fatalf("no se esperaba error, pero se obtuvo: %v", err)
	}

	if libro.ID() != 1 {
		t.Errorf("se esperaba ID 1, se obtuvo %d", libro.ID())
	}

	if libro.Titulo() != "Introduccion a Go" {
		t.Errorf("titulo incorrecto: %s", libro.Titulo())
	}

	if !libro.Disponible() {
		t.Errorf("un libro nuevo debe estar disponible por defecto")
	}
}

func TestNuevoLibroConTituloVacio(t *testing.T) {
	_, err := NuevoLibro(1, "", "Alan Donovan", "Programacion", 2016)

	if err == nil {
		t.Fatal("se esperaba error cuando el titulo esta vacio")
	}
}

func TestCambiarDisponibilidad(t *testing.T) {
	libro, err := NuevoLibro(1, "Clean Code", "Robert C. Martin", "Programacion", 2008)

	if err != nil {
		t.Fatalf("no se esperaba error, pero se obtuvo: %v", err)
	}

	libro.CambiarDisponibilidad(false)

	if libro.Disponible() {
		t.Error("se esperaba que el libro no este disponible")
	}
}

func TestSetAnioInvalido(t *testing.T) {
	libro, err := NuevoLibro(1, "Clean Code", "Robert C. Martin", "Programacion", 2008)

	if err != nil {
		t.Fatalf("no se esperaba error, pero se obtuvo: %v", err)
	}

	err = libro.SetAnio(0)

	if err == nil {
		t.Fatal("se esperaba error para un anio invalido")
	}
}
