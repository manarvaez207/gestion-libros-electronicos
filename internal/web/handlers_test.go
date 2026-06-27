package web

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"sync"
	"testing"
)

func nuevaAppDePrueba(t *testing.T) *App {
	t.Helper()

	rutaDatos := filepath.Join(t.TempDir(), "libros.json")
	return NuevoApp(rutaDatos)
}

func TestGETLibros(t *testing.T) {
	app := nuevaAppDePrueba(t)

	req := httptest.NewRequest(http.MethodGet, "/libros", nil)
	rec := httptest.NewRecorder()

	app.Rutas().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("se esperaba estado 200, se obtuvo %d", rec.Code)
	}

	if rec.Body.Len() == 0 {
		t.Fatal("se esperaba una respuesta JSON con libros")
	}
}

func TestPOSTCrearLibro(t *testing.T) {
	app := nuevaAppDePrueba(t)

	cuerpo := []byte(`{
		"titulo": "Clean Code",
		"autor": "Robert C. Martin",
		"categoria": "Programacion",
		"anio": 2008,
		"disponible": true
	}`)

	req := httptest.NewRequest(http.MethodPost, "/libros", bytes.NewBuffer(cuerpo))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	app.Rutas().ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("se esperaba estado 201, se obtuvo %d. Respuesta: %s", rec.Code, rec.Body.String())
	}
}

func TestPUTActualizarLibro(t *testing.T) {
	app := nuevaAppDePrueba(t)

	cuerpo := []byte(`{
		"titulo": "Introduccion a Go Actualizado",
		"autor": "Alan Donovan",
		"categoria": "Programacion",
		"anio": 2017,
		"disponible": true
	}`)

	req := httptest.NewRequest(http.MethodPut, "/libros/1", bytes.NewBuffer(cuerpo))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	app.Rutas().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("se esperaba estado 200, se obtuvo %d. Respuesta: %s", rec.Code, rec.Body.String())
	}
}

func TestPATCHCambiarDisponibilidad(t *testing.T) {
	app := nuevaAppDePrueba(t)

	cuerpo := []byte(`{
		"disponible": false
	}`)

	req := httptest.NewRequest(http.MethodPatch, "/libros/1/disponibilidad", bytes.NewBuffer(cuerpo))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	app.Rutas().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("se esperaba estado 200, se obtuvo %d. Respuesta: %s", rec.Code, rec.Body.String())
	}
}

func TestDELETEEliminarLibro(t *testing.T) {
	app := nuevaAppDePrueba(t)

	req := httptest.NewRequest(http.MethodDelete, "/libros/1", nil)
	rec := httptest.NewRecorder()

	app.Rutas().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("se esperaba estado 200, se obtuvo %d. Respuesta: %s", rec.Code, rec.Body.String())
	}
}

func TestReporteGeneral(t *testing.T) {
	app := nuevaAppDePrueba(t)

	req := httptest.NewRequest(http.MethodGet, "/reportes/general", nil)
	rec := httptest.NewRecorder()

	app.Rutas().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("se esperaba estado 200, se obtuvo %d", rec.Code)
	}

	if rec.Body.Len() == 0 {
		t.Fatal("se esperaba respuesta JSON del reporte general")
	}
}

func TestSolicitudesConcurrentes(t *testing.T) {
	app := nuevaAppDePrueba(t)

	var wg sync.WaitGroup
	totalSolicitudes := 20

	for i := 0; i < totalSolicitudes; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			req := httptest.NewRequest(http.MethodGet, "/libros", nil)
			rec := httptest.NewRecorder()

			app.Rutas().ServeHTTP(rec, req)

			if rec.Code != http.StatusOK {
				t.Errorf("se esperaba estado 200, se obtuvo %d", rec.Code)
			}
		}()
	}

	wg.Wait()
}
