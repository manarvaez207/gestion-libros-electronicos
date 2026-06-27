package web

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"sistema-libros/internal/categorias"
	"sistema-libros/internal/libros"
)

// App representa la aplicación web del sistema.
// Mantiene en memoria los libros y categorías, y además guarda los cambios en JSON.
type App struct {
	libros     []libros.Libro
	categorias []categorias.Categoria
	rutaDatos  string
	mu         sync.Mutex
}

// libroDTO permite serializar y deserializar libros mediante JSON.
// Se usa porque el struct Libro original tiene atributos privados por encapsulación.
type libroDTO struct {
	ID         int    `json:"id"`
	Titulo     string `json:"titulo"`
	Autor      string `json:"autor"`
	Categoria  string `json:"categoria"`
	Anio       int    `json:"anio"`
	Disponible bool   `json:"disponible"`
}

// libroRequest representa los datos recibidos al crear o actualizar un libro.
type libroRequest struct {
	Titulo     string `json:"titulo"`
	Autor      string `json:"autor"`
	Categoria  string `json:"categoria"`
	Anio       int    `json:"anio"`
	Disponible *bool  `json:"disponible,omitempty"`
}

// disponibilidadRequest representa el cuerpo JSON para cambiar disponibilidad.
type disponibilidadRequest struct {
	Disponible bool `json:"disponible"`
}

// categoriaDTO permite responder categorías en formato JSON.
type categoriaDTO struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
}

// categoriaRequest representa el cuerpo JSON para crear una categoría.
type categoriaRequest struct {
	Nombre string `json:"nombre"`
}

// reporteGeneralDTO representa el reporte general del sistema.
type reporteGeneralDTO struct {
	TotalLibros         int `json:"total_libros"`
	LibrosDisponibles   int `json:"libros_disponibles"`
	LibrosNoDisponibles int `json:"libros_no_disponibles"`
}

// reporteCategoriaDTO representa el reporte de libros por categoría.
type reporteCategoriaDTO struct {
	Categoria   string `json:"categoria"`
	TotalLibros int    `json:"total_libros"`
}

// respuestaError permite enviar errores en formato JSON.
type respuestaError struct {
	Error string `json:"error"`
}

// NuevoApp crea la aplicación web, carga los datos desde JSON y prepara datos iniciales.
func NuevoApp(rutaDatos string) *App {
	app := &App{
		rutaDatos:  rutaDatos,
		libros:     []libros.Libro{},
		categorias: []categorias.Categoria{},
	}

	_ = app.cargarDesdeJSON()

	// Si no existen libros, se crea un libro inicial para facilitar la demostración.
	if len(app.libros) == 0 {
		libroInicial, err := libros.NuevoLibro(
			1,
			"Introducción a Go",
			"Alan Donovan",
			"Programación",
			2016,
		)

		if err == nil {
			app.libros = append(app.libros, libroInicial)
			_ = app.guardarEnJSON()
		}
	}

	app.sincronizarCategoriasDesdeLibros()
	return app
}

// Rutas registra todos los servicios web del sistema.
func (app *App) Rutas() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.handleInicio)

	// Servicios web de libros
	mux.HandleFunc("/libros", app.handleLibros)
	mux.HandleFunc("/libros/", app.handleLibroPorID)
	mux.HandleFunc("/libros/buscar", app.handleBuscarPorTitulo)

	// Servicios web de categorías
	mux.HandleFunc("/categorias", app.handleCategorias)

	// Servicios web de reportes
	mux.HandleFunc("/reportes/general", app.handleReporteGeneral)
	mux.HandleFunc("/reportes/categoria", app.handleReportePorCategoria)

	return mux
}

// handleInicio muestra información general de la API.
func (app *App) handleInicio(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		responderError(w, http.StatusNotFound, "ruta no encontrada")
		return
	}

	respuesta := map[string]interface{}{
		"mensaje": "API Web del Sistema de Gestión de Libros Electrónicos",
		"version": "Proyecto final",
		"servicios_disponibles": []string{
			"GET /libros",
			"GET /libros/{id}",
			"POST /libros",
			"PUT /libros/{id}",
			"DELETE /libros/{id}",
			"PATCH /libros/{id}/disponibilidad",
			"GET /libros/buscar?titulo=go",
			"GET /categorias",
			"POST /categorias",
			"GET /reportes/general",
			"GET /reportes/categoria?nombre=Programación",
		},
	}

	responderJSON(w, http.StatusOK, respuesta)
}

// handleLibros gestiona GET /libros y POST /libros.
func (app *App) handleLibros(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/libros" {
		responderError(w, http.StatusNotFound, "ruta no encontrada")
		return
	}

	switch r.Method {
	case http.MethodGet:
		app.listarLibros(w, r)
	case http.MethodPost:
		app.crearLibro(w, r)
	default:
		responderError(w, http.StatusMethodNotAllowed, "método no permitido")
	}
}

// listarLibros responde todos los libros registrados.
func (app *App) listarLibros(w http.ResponseWriter, r *http.Request) {
	app.mu.Lock()
	defer app.mu.Unlock()

	responderJSON(w, http.StatusOK, app.convertirLibrosADTO(app.libros))
}

// crearLibro registra un nuevo libro usando JSON.
func (app *App) crearLibro(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req libroRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	app.mu.Lock()
	defer app.mu.Unlock()

	nuevoLibro, err := libros.NuevoLibro(
		app.siguienteIDLibro(),
		req.Titulo,
		req.Autor,
		req.Categoria,
		req.Anio,
	)

	if err != nil {
		responderError(w, http.StatusBadRequest, err.Error())
		return
	}

	if req.Disponible != nil {
		nuevoLibro.CambiarDisponibilidad(*req.Disponible)
	}

	app.libros = append(app.libros, nuevoLibro)
	app.crearCategoriaSiNoExiste(req.Categoria)

	if err := app.guardarEnJSON(); err != nil {
		responderError(w, http.StatusInternalServerError, "no se pudo guardar el libro")
		return
	}

	responderJSON(w, http.StatusCreated, convertirLibroADTO(nuevoLibro))
}

// handleLibroPorID gestiona rutas relacionadas con un libro específico.
func (app *App) handleLibroPorID(w http.ResponseWriter, r *http.Request) {
	partes := strings.Split(strings.TrimPrefix(r.URL.Path, "/libros/"), "/")

	if len(partes) == 0 || partes[0] == "" {
		responderError(w, http.StatusBadRequest, "ID de libro no válido")
		return
	}

	id, err := strconv.Atoi(partes[0])
	if err != nil {
		responderError(w, http.StatusBadRequest, "ID de libro no válido")
		return
	}

	// Ruta: PATCH /libros/{id}/disponibilidad
	if len(partes) == 2 && partes[1] == "disponibilidad" {
		if r.Method == http.MethodPatch {
			app.cambiarDisponibilidad(w, r, id)
			return
		}

		responderError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}

	if len(partes) > 1 {
		responderError(w, http.StatusNotFound, "ruta no encontrada")
		return
	}

	switch r.Method {
	case http.MethodGet:
		app.obtenerLibroPorID(w, r, id)
	case http.MethodPut:
		app.actualizarLibro(w, r, id)
	case http.MethodDelete:
		app.eliminarLibro(w, r, id)
	default:
		responderError(w, http.StatusMethodNotAllowed, "método no permitido")
	}
}

// obtenerLibroPorID responde un libro específico.
func (app *App) obtenerLibroPorID(w http.ResponseWriter, r *http.Request, id int) {
	app.mu.Lock()
	defer app.mu.Unlock()

	indice := app.buscarIndiceLibro(id)
	if indice == -1 {
		responderError(w, http.StatusNotFound, "libro no encontrado")
		return
	}

	responderJSON(w, http.StatusOK, convertirLibroADTO(app.libros[indice]))
}

// actualizarLibro actualiza completamente la información de un libro.
func (app *App) actualizarLibro(w http.ResponseWriter, r *http.Request, id int) {
	defer r.Body.Close()

	var req libroRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	app.mu.Lock()
	defer app.mu.Unlock()

	indice := app.buscarIndiceLibro(id)
	if indice == -1 {
		responderError(w, http.StatusNotFound, "libro no encontrado")
		return
	}

	libroActualizado, err := libros.NuevoLibro(
		id,
		req.Titulo,
		req.Autor,
		req.Categoria,
		req.Anio,
	)

	if err != nil {
		responderError(w, http.StatusBadRequest, err.Error())
		return
	}

	if req.Disponible != nil {
		libroActualizado.CambiarDisponibilidad(*req.Disponible)
	} else {
		libroActualizado.CambiarDisponibilidad(app.libros[indice].Disponible())
	}

	app.libros[indice] = libroActualizado
	app.crearCategoriaSiNoExiste(req.Categoria)

	if err := app.guardarEnJSON(); err != nil {
		responderError(w, http.StatusInternalServerError, "no se pudo actualizar el libro")
		return
	}

	responderJSON(w, http.StatusOK, convertirLibroADTO(libroActualizado))
}

// eliminarLibro elimina un libro por ID.
func (app *App) eliminarLibro(w http.ResponseWriter, r *http.Request, id int) {
	app.mu.Lock()
	defer app.mu.Unlock()

	indice := app.buscarIndiceLibro(id)
	if indice == -1 {
		responderError(w, http.StatusNotFound, "libro no encontrado")
		return
	}

	libroEliminado := app.libros[indice]
	app.libros = append(app.libros[:indice], app.libros[indice+1:]...)

	if err := app.guardarEnJSON(); err != nil {
		responderError(w, http.StatusInternalServerError, "no se pudo eliminar el libro")
		return
	}

	responderJSON(w, http.StatusOK, map[string]interface{}{
		"mensaje": "libro eliminado correctamente",
		"libro":   convertirLibroADTO(libroEliminado),
	})
}

// cambiarDisponibilidad actualiza únicamente el estado disponible/no disponible.
func (app *App) cambiarDisponibilidad(w http.ResponseWriter, r *http.Request, id int) {
	defer r.Body.Close()

	var req disponibilidadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	app.mu.Lock()
	defer app.mu.Unlock()

	indice := app.buscarIndiceLibro(id)
	if indice == -1 {
		responderError(w, http.StatusNotFound, "libro no encontrado")
		return
	}

	app.libros[indice].CambiarDisponibilidad(req.Disponible)

	if err := app.guardarEnJSON(); err != nil {
		responderError(w, http.StatusInternalServerError, "no se pudo actualizar la disponibilidad")
		return
	}

	responderJSON(w, http.StatusOK, convertirLibroADTO(app.libros[indice]))
}

// handleBuscarPorTitulo busca libros mediante el parámetro titulo.
func (app *App) handleBuscarPorTitulo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		responderError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}

	titulo := strings.TrimSpace(r.URL.Query().Get("titulo"))
	if titulo == "" {
		responderError(w, http.StatusBadRequest, "debe enviar el parámetro titulo")
		return
	}

	app.mu.Lock()
	defer app.mu.Unlock()

	resultados := []libros.Libro{}
	titulo = strings.ToLower(titulo)

	for _, libro := range app.libros {
		if strings.Contains(strings.ToLower(libro.Titulo()), titulo) {
			resultados = append(resultados, libro)
		}
	}

	responderJSON(w, http.StatusOK, app.convertirLibrosADTO(resultados))
}

// handleCategorias gestiona GET /categorias y POST /categorias.
func (app *App) handleCategorias(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/categorias" {
		responderError(w, http.StatusNotFound, "ruta no encontrada")
		return
	}

	switch r.Method {
	case http.MethodGet:
		app.listarCategorias(w, r)
	case http.MethodPost:
		app.crearCategoria(w, r)
	default:
		responderError(w, http.StatusMethodNotAllowed, "método no permitido")
	}
}

// listarCategorias responde las categorías registradas.
func (app *App) listarCategorias(w http.ResponseWriter, r *http.Request) {
	app.mu.Lock()
	defer app.mu.Unlock()

	respuesta := []categoriaDTO{}

	for _, categoria := range app.categorias {
		respuesta = append(respuesta, categoriaDTO{
			ID:     categoria.ID(),
			Nombre: categoria.Nombre(),
		})
	}

	responderJSON(w, http.StatusOK, respuesta)
}

// crearCategoria registra una nueva categoría desde JSON.
func (app *App) crearCategoria(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req categoriaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	app.mu.Lock()
	defer app.mu.Unlock()

	if app.existeCategoria(req.Nombre) {
		responderError(w, http.StatusBadRequest, "la categoría ya existe")
		return
	}

	nuevaCategoria, err := categorias.NuevaCategoria(app.siguienteIDCategoria(), req.Nombre)
	if err != nil {
		responderError(w, http.StatusBadRequest, err.Error())
		return
	}

	app.categorias = append(app.categorias, nuevaCategoria)

	responderJSON(w, http.StatusCreated, categoriaDTO{
		ID:     nuevaCategoria.ID(),
		Nombre: nuevaCategoria.Nombre(),
	})
}

// handleReporteGeneral genera un resumen general del catálogo.
func (app *App) handleReporteGeneral(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		responderError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}

	app.mu.Lock()
	defer app.mu.Unlock()

	disponibles := 0
	noDisponibles := 0

	for _, libro := range app.libros {
		if libro.Disponible() {
			disponibles++
		} else {
			noDisponibles++
		}
	}

	reporte := reporteGeneralDTO{
		TotalLibros:         len(app.libros),
		LibrosDisponibles:   disponibles,
		LibrosNoDisponibles: noDisponibles,
	}

	responderJSON(w, http.StatusOK, reporte)
}

// handleReportePorCategoria genera un reporte filtrado por categoría.
func (app *App) handleReportePorCategoria(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		responderError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}

	nombre := strings.TrimSpace(r.URL.Query().Get("nombre"))
	if nombre == "" {
		responderError(w, http.StatusBadRequest, "debe enviar el parámetro nombre")
		return
	}

	app.mu.Lock()
	defer app.mu.Unlock()

	total := 0

	for _, libro := range app.libros {
		if strings.EqualFold(libro.Categoria(), nombre) {
			total++
		}
	}

	reporte := reporteCategoriaDTO{
		Categoria:   nombre,
		TotalLibros: total,
	}

	responderJSON(w, http.StatusOK, reporte)
}

// cargarDesdeJSON carga los libros almacenados en el archivo JSON.
func (app *App) cargarDesdeJSON() error {
	contenido, err := os.ReadFile(app.rutaDatos)
	if err != nil {
		return err
	}

	var librosJSON []libroDTO
	if err := json.Unmarshal(contenido, &librosJSON); err != nil {
		return err
	}

	for _, item := range librosJSON {
		libro, err := libros.NuevoLibro(
			item.ID,
			item.Titulo,
			item.Autor,
			item.Categoria,
			item.Anio,
		)

		if err != nil {
			continue
		}

		libro.CambiarDisponibilidad(item.Disponible)
		app.libros = append(app.libros, libro)
	}

	return nil
}

// guardarEnJSON guarda los libros en el archivo JSON.
func (app *App) guardarEnJSON() error {
	librosJSON := app.convertirLibrosADTO(app.libros)

	contenido, err := json.MarshalIndent(librosJSON, "", "  ")
	if err != nil {
		return err
	}

	carpeta := filepath.Dir(app.rutaDatos)
	if err := os.MkdirAll(carpeta, 0755); err != nil {
		return err
	}

	return os.WriteFile(app.rutaDatos, contenido, 0644)
}

// sincronizarCategoriasDesdeLibros genera categorías según los libros existentes.
func (app *App) sincronizarCategoriasDesdeLibros() {
	for _, libro := range app.libros {
		app.crearCategoriaSiNoExiste(libro.Categoria())
	}
}

// crearCategoriaSiNoExiste agrega una categoría si aún no está registrada.
func (app *App) crearCategoriaSiNoExiste(nombre string) {
	nombre = strings.TrimSpace(nombre)
	if nombre == "" || app.existeCategoria(nombre) {
		return
	}

	nuevaCategoria, err := categorias.NuevaCategoria(app.siguienteIDCategoria(), nombre)
	if err == nil {
		app.categorias = append(app.categorias, nuevaCategoria)
	}
}

// existeCategoria verifica si una categoría ya está registrada.
func (app *App) existeCategoria(nombre string) bool {
	for _, categoria := range app.categorias {
		if strings.EqualFold(categoria.Nombre(), nombre) {
			return true
		}
	}

	return false
}

// buscarIndiceLibro obtiene la posición de un libro por ID.
func (app *App) buscarIndiceLibro(id int) int {
	for indice, libro := range app.libros {
		if libro.ID() == id {
			return indice
		}
	}

	return -1
}

// siguienteIDLibro genera el siguiente ID disponible para libros.
func (app *App) siguienteIDLibro() int {
	mayor := 0

	for _, libro := range app.libros {
		if libro.ID() > mayor {
			mayor = libro.ID()
		}
	}

	return mayor + 1
}

// siguienteIDCategoria genera el siguiente ID disponible para categorías.
func (app *App) siguienteIDCategoria() int {
	mayor := 0

	for _, categoria := range app.categorias {
		if categoria.ID() > mayor {
			mayor = categoria.ID()
		}
	}

	return mayor + 1
}

// convertirLibrosADTO convierte una lista de libros encapsulados a una lista serializable.
func (app *App) convertirLibrosADTO(lista []libros.Libro) []libroDTO {
	respuesta := []libroDTO{}

	for _, libro := range lista {
		respuesta = append(respuesta, convertirLibroADTO(libro))
	}

	return respuesta
}

// convertirLibroADTO convierte un libro a una estructura compatible con JSON.
func convertirLibroADTO(libro libros.Libro) libroDTO {
	return libroDTO{
		ID:         libro.ID(),
		Titulo:     libro.Titulo(),
		Autor:      libro.Autor(),
		Categoria:  libro.Categoria(),
		Anio:       libro.Anio(),
		Disponible: libro.Disponible(),
	}
}

// responderJSON envía una respuesta en formato JSON.
func responderJSON(w http.ResponseWriter, estado int, datos interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(estado)
	_ = json.NewEncoder(w).Encode(datos)
}

// responderError envía un mensaje de error en formato JSON.
func responderError(w http.ResponseWriter, estado int, mensaje string) {
	responderJSON(w, estado, respuestaError{
		Error: mensaje,
	})
}
