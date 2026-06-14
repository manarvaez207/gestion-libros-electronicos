package main

import (
	"fmt"
	"log"

	"sistema-libros/internal/busquedas"
	"sistema-libros/internal/categorias"
	"sistema-libros/internal/libros"
	"sistema-libros/internal/persistencia"
	"sistema-libros/internal/reportes"
)

func main() {
	fmt.Println("Sistema de Gestión de Libros Electrónicos")
	fmt.Println("----------------------------------------")

	categoria1, err := categorias.NuevaCategoria(1, "Programación")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Categoría registrada correctamente:")
	fmt.Println("ID:", categoria1.ID(), "| Nombre:", categoria1.Nombre())

	libro1, err := libros.NuevoLibro(
		1,
		"Introducción a Go",
		"Alan Donovan",
		categoria1.Nombre(),
		2016,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nLibro registrado correctamente:")
	fmt.Println(libro1.Descripcion())

	fmt.Println("\nActualizando disponibilidad del libro...")
	libro1.CambiarDisponibilidad(false)

	fmt.Println(libro1.Descripcion())

	catalogoBusquedas := []busquedas.LibroBuscable{
		libro1,
	}

	fmt.Println("\nBuscando libro por título...")
	resultadosTitulo, err := busquedas.BuscarPorTitulo(catalogoBusquedas, "go")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, libro := range resultadosTitulo {
			fmt.Println(libro.Descripcion())
		}
	}

	fmt.Println("\nBuscando libro por categoría...")
	resultadosCategoria, err := busquedas.BuscarPorCategoria(catalogoBusquedas, "programación")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, libro := range resultadosCategoria {
			fmt.Println(libro.Descripcion())
		}
	}

	catalogoLibros := []libros.Libro{
		libro1,
	}

	fmt.Println("\nGuardando libros en archivo JSON...")
	err = persistencia.GuardarLibros("data/libros.json", catalogoLibros)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Libros guardados correctamente en data/libros.json")

	fmt.Println("\nCargando libros desde archivo JSON...")
	librosCargados, err := persistencia.CargarLibros("data/libros.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Libros cargados correctamente:")
	for _, libro := range librosCargados {
		fmt.Println(libro.Descripcion())
	}

	catalogoReportes := []reportes.LibroReporte{
		libro1,
	}

	fmt.Println("\nGenerando reporte general...")
	fmt.Println(reportes.GenerarResumen(catalogoReportes))

	fmt.Println("\nGenerando reporte por categoría...")
	totalCategoria, err := reportes.ContarPorCategoria(catalogoReportes, "programación")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Total de libros en la categoría Programación:", totalCategoria)
	}
}
