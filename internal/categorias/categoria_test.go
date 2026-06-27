package categorias

import "testing"

func TestNuevaCategoriaValida(t *testing.T) {
	categoria, err := NuevaCategoria(1, "Programacion")

	if err != nil {
		t.Fatalf("no se esperaba error, pero se obtuvo: %v", err)
	}

	if categoria.ID() != 1 {
		t.Errorf("se esperaba ID 1, se obtuvo %d", categoria.ID())
	}

	if categoria.Nombre() != "Programacion" {
		t.Errorf("nombre incorrecto: %s", categoria.Nombre())
	}
}

func TestNuevaCategoriaConNombreVacio(t *testing.T) {
	_, err := NuevaCategoria(1, "")

	if err == nil {
		t.Fatal("se esperaba error cuando el nombre de la categoria esta vacio")
	}
}

func TestSetNombreCategoria(t *testing.T) {
	categoria, err := NuevaCategoria(1, "Programacion")

	if err != nil {
		t.Fatalf("no se esperaba error, pero se obtuvo: %v", err)
	}

	err = categoria.SetNombre("Tecnologia")

	if err != nil {
		t.Fatalf("no se esperaba error al cambiar el nombre, pero se obtuvo: %v", err)
	}

	if categoria.Nombre() != "Tecnologia" {
		t.Errorf("se esperaba Tecnologia, se obtuvo %s", categoria.Nombre())
	}
}
