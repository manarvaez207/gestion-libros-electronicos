package categorias

import "errors"

// Categoria representa una categoría dentro del sistema de libros electrónicos.
// Sus atributos están encapsulados para evitar modificaciones directas.
type Categoria struct {
	id     int
	nombre string
}

// NuevaCategoria permite crear una categoría validando sus datos.
func NuevaCategoria(id int, nombre string) (Categoria, error) {
	categoria := Categoria{
		id: id,
	}

	if err := categoria.SetNombre(nombre); err != nil {
		return Categoria{}, err
	}

	return categoria, nil
}

// ID devuelve el identificador de la categoría.
func (c Categoria) ID() int {
	return c.id
}

// Nombre devuelve el nombre de la categoría.
func (c Categoria) Nombre() string {
	return c.nombre
}

// SetNombre asigna un nombre a la categoría validando que no esté vacío.
func (c *Categoria) SetNombre(nombre string) error {
	if nombre == "" {
		return errors.New("el nombre de la categoría no puede estar vacío")
	}

	c.nombre = nombre
	return nil
}
