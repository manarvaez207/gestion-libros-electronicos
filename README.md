# Sistema de Gestión de Libros Electrónicos

## Aprendizaje Autónomo 2: Desarrollo del Sistema de Gestión

Este proyecto corresponde al desarrollo de un sistema de gestión de libros electrónicos implementado en el lenguaje de programación **Go (Golang)**. El sistema permite registrar libros, gestionar categorías, realizar búsquedas, modificar la disponibilidad de los libros, guardar información en un archivo JSON y generar reportes básicos del catálogo.

El trabajo continúa la planificación realizada en el Aprendizaje Autónomo 1 y aplica los contenidos estudiados en las unidades de la materia, especialmente estructuras de datos, objetos en Go, encapsulación, manejo de errores e interfaces.

---

## Objetivo del proyecto

Desarrollar un sistema de gestión de libros electrónicos en Go, aplicando programación modular, estructuras de datos, encapsulación, manejo de errores, interfaces y persistencia de datos mediante archivos JSON.

---

## Funcionalidades implementadas

El sistema desarrollado incluye las siguientes funcionalidades:

- Registro de categorías.
- Registro de libros electrónicos.
- Validación de datos al crear libros y categorías.
- Encapsulación de atributos mediante campos privados.
- Uso de métodos getter para consultar información.
- Uso de métodos setter para modificar información de forma controlada.
- Cambio de disponibilidad de libros.
- Búsqueda de libros por título.
- Búsqueda de libros por categoría.
- Persistencia de libros en archivo JSON.
- Carga de libros desde archivo JSON.
- Generación de reporte general del sistema.
- Generación de reporte por categoría.

---

## Estructura del proyecto

La estructura del proyecto está organizada de forma modular:

```text
Universidad el ProyectoGo/
│
├── cmd/
│   └── main.go
│
├── data/
│   └── libros.json
│
├── docs/
│
├── internal/
│   ├── busquedas/
│   │   └── busquedas.go
│   │
│   ├── categorias/
│   │   └── categoria.go
│   │
│   ├── libros/
│   │   └── libro.go
│   │
│   ├── persistencia/
│   │   └── persistencia.go
│   │
│   └── reportes/
│       └── reportes.go
│
├── go.mod
└── README.md

Descripción de los paquetes
cmd

Contiene el archivo principal main.go, desde donde se ejecuta el sistema. En este archivo se integran los módulos del proyecto para demostrar el funcionamiento general del sistema.

internal/libros

Contiene la estructura principal Libro, sus métodos getter, métodos setter, validaciones y funciones relacionadas con la creación y modificación de libros electrónicos.

Este paquete aplica encapsulación porque los atributos del struct se mantienen privados y solo se accede a ellos mediante métodos.

internal/categorias

Contiene la estructura Categoria, utilizada para registrar y validar categorías dentro del sistema.

internal/busquedas

Contiene funciones para buscar libros por título y por categoría. Este paquete utiliza una interfaz llamada LibroBuscable, lo que permite aplicar polimorfismo y desacoplar la búsqueda del tipo concreto del libro.

internal/persistencia

Contiene funciones para guardar y cargar libros desde un archivo JSON. Esto permite que la información del sistema pueda conservarse fuera de la ejecución del programa.

internal/reportes

Contiene funciones para generar reportes básicos del sistema, como el total de libros registrados, libros disponibles, libros no disponibles y cantidad de libros por categoría.

data

Contiene el archivo libros.json, donde se almacenan los datos guardados por el sistema.

docs

Carpeta reservada para documentación del proyecto, evidencias o material adicional.

Conceptos aplicados

En el desarrollo del sistema se aplicaron los siguientes conceptos estudiados en clase:

Estructuras de datos

Se utilizaron slices para manejar colecciones de libros dentro del sistema. Esto permite almacenar varios libros y recorrerlos mediante ciclos.

Structs

Se utilizaron structs para representar entidades del sistema, como libros y categorías.

Encapsulación

Los atributos de los structs fueron definidos con letra minúscula para mantenerlos privados dentro de su paquete. Para acceder o modificar estos datos, se crearon métodos getter y setter.

Métodos getter

Permiten obtener información de un objeto sin acceder directamente a sus atributos internos.

Métodos setter

Permiten modificar los atributos de forma controlada, incluyendo validaciones antes de asignar nuevos valores.

Manejo de errores

Se implementó manejo de errores con el tipo error, utilizando validaciones para evitar datos vacíos, años inválidos o búsquedas sin resultados.

Interfaces

Se utilizó la interfaz LibroBuscable para definir los métodos necesarios para realizar búsquedas sobre los libros. Esto permite trabajar de forma más flexible y aplicar polimorfismo.

Persistencia con JSON

Se utilizó el formato JSON para guardar y cargar información del sistema desde un archivo externo.

Programación modular

El sistema fue dividido en paquetes para mejorar la organización, legibilidad y mantenimiento del código.

Paquetes utilizados

El proyecto utiliza únicamente paquetes estándar de Go:

fmt
log
errors
strings
encoding/json
os

No se utilizaron paquetes de terceros.

Ejecución del proyecto

Para ejecutar el sistema, se debe abrir la terminal en la carpeta raíz del proyecto y ejecutar el siguiente comando:

go run ./cmd
Formateo del código

Para dar formato automático al código fuente, se utilizó el comando:

go fmt ./...

Este comando organiza todos los archivos .go del proyecto siguiendo el formato estándar de Go.

Verificación del proyecto

Para verificar que el proyecto no tenga errores de compilación, se puede ejecutar:

go test ./...

Aunque el proyecto no incluye archivos de prueba específicos, este comando permite comprobar que todos los paquetes compilan correctamente.

Sistema de Gestión de Libros Electrónicos
----------------------------------------

Categoría registrada correctamente:
ID: 1 | Nombre: Programación

Libro registrado correctamente:
ID: 1 | Título: Introducción a Go | Autor: Alan Donovan | Categoría: Programación | Año: 2016 | Estado: Disponible

Actualizando disponibilidad del libro...
ID: 1 | Título: Introducción a Go | Autor: Alan Donovan | Categoría: Programación | Año: 2016 | Estado: No disponible

Buscando libro por título...
ID: 1 | Título: Introducción a Go | Autor: Alan Donovan | Categoría: Programación | Año: 2016 | Estado: No disponible

Buscando libro por categoría...
ID: 1 | Título: Introducción a Go | Autor: Alan Donovan | Categoría: Programación | Año: 2016 | Estado: No disponible

Guardando libros en archivo JSON...
Libros guardados correctamente en data/libros.json

Cargando libros desde archivo JSON...
Libros cargados correctamente:
ID: 1 | Título: Introducción a Go | Autor: Alan Donovan | Categoría: Programación | Año: 2016 | Estado: No disponible

Generando reporte general...
Resumen del sistema:
Total de libros: 1
Libros disponibles: 0
Libros no disponibles: 1

Generando reporte por categoría...
Total de libros en la categoría Programación: 1

Archivo JSON generado

El sistema guarda los libros en el archivo data/libros.json. Un ejemplo del contenido generado es:

[
  {
    "id": 1,
    "titulo": "Introducción a Go",
    "autor": "Alan Donovan",
    "categoria": "Programación",
    "anio": 2016,
    "disponible": false
  }
]
Relación con los temas estudiados

Este proyecto aplica los temas revisados durante las unidades de la asignatura:

Unidad 1
Variables.
Tipos de datos.
Operadores.
Condicionales.
Bucles.
Funciones.
Paquetes.
Módulos.
Unidad 2
Arrays.
Slices.
Maps.
Structs.
Métodos.
Funciones.
Constructores mediante funciones.
Unidad 3
Encapsulación.
Métodos setter.
Métodos getter.
Manejo de errores.
Interfaces.
Implementación de interfaces.
Polimorfismo mediante interfaces.
Conclusión

El sistema de gestión de libros electrónicos desarrollado en Go permite aplicar de manera práctica los contenidos estudiados en la asignatura. A través de una estructura modular, el proyecto integra structs, métodos, encapsulación, manejo de errores, interfaces, persistencia en JSON y reportes básicos.

Además, el desarrollo del sistema demuestra cómo Go permite construir aplicaciones organizadas, legibles y mantenibles mediante el uso adecuado de paquetes, estructuras de datos y programación orientada a objetos de forma simplificada.

Autor

Desarrollado por: Matthew Adrian Narváez Paredes
Proyecto individual
Materia: Programación en Go