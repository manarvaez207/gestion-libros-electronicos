# Aprendizaje Autónomo 2: Desarrollo del Sistema de Gestión

## Sistema de Gestión de Libros Electrónicos

**Estudiante:** Matthew Narváez
**Lenguaje de programación:** Go / Golang
**Repositorio:** https://github.com/manarvaez207/gestion-libros-electronicos

---

## 1. Introducción

El presente documento describe el avance realizado en el proyecto correspondiente al Aprendizaje Autónomo 2: Desarrollo del Sistema de Gestión. El sistema seleccionado es un Sistema de Gestión de Libros Electrónicos, desarrollado en el lenguaje de programación Go.

Este avance continúa la planificación realizada en el Aprendizaje Autónomo 1, donde se definieron el objetivo del sistema, los módulos principales y las funcionalidades esperadas. En esta segunda etapa se inició la codificación del sistema, aplicando los contenidos estudiados en las unidades de la materia, especialmente estructuras de datos, objetos en Go, encapsulación, manejo de errores e interfaces.

---

## 2. Objetivo del proyecto

Desarrollar una versión funcional inicial de un Sistema de Gestión de Libros Electrónicos en Go, que permita registrar libros, gestionar categorías, realizar búsquedas, modificar la disponibilidad de los libros, guardar información en archivos JSON y generar reportes básicos del catálogo.

---

## 3. Continuación del plan

El proyecto mantiene la línea definida en el Aprendizaje Autónomo 1. En la etapa anterior se seleccionó el sistema de gestión de libros electrónicos y se planteó su estructura general. En esta nueva etapa se inició la implementación del sistema mediante código fuente organizado en paquetes.

El desarrollo actual representa un avance significativo porque el sistema ya permite ejecutar funcionalidades reales desde la terminal, guardar datos en un archivo JSON y mostrar reportes básicos. Además, la estructura modular facilita el mantenimiento y la ampliación del sistema en futuras versiones.

---

## 4. Estructura del proyecto

La estructura del proyecto se organizó de la siguiente manera:

* `cmd/main.go`: archivo principal desde donde se ejecuta el sistema.
* `data/libros.json`: archivo utilizado para guardar la información de los libros.
* `internal/libros/libro.go`: paquete encargado de gestionar la entidad Libro.
* `internal/categorias/categoria.go`: paquete encargado de gestionar categorías.
* `internal/busquedas/busquedas.go`: paquete encargado de realizar búsquedas por título y categoría.
* `internal/persistencia/persistencia.go`: paquete encargado de guardar y cargar datos en formato JSON.
* `internal/reportes/reportes.go`: paquete encargado de generar reportes del sistema.
* `README.md`: documentación general del proyecto.
* `go.mod`: archivo del módulo de Go.

Esta organización permite separar responsabilidades y mantener el código más claro, modular y fácil de comprender.

---

## 5. Funcionalidades implementadas

El sistema desarrollado incluye las siguientes funcionalidades:

* Registro de categorías.
* Registro de libros electrónicos.
* Validación de datos al crear libros y categorías.
* Encapsulación de atributos mediante campos privados.
* Uso de métodos getter para consultar información.
* Uso de métodos setter para modificar datos de forma controlada.
* Cambio de disponibilidad de libros.
* Búsqueda de libros por título.
* Búsqueda de libros por categoría.
* Persistencia de libros en archivo JSON.
* Carga de libros desde archivo JSON.
* Generación de reporte general del sistema.
* Generación de reporte por categoría.

---

## 6. Aplicación de estructuras de datos y objetos en Go

En el proyecto se aplican estructuras de datos como slices y structs.

Los slices se utilizan para manejar listas de libros dentro del sistema. Esto permite recorrer, buscar, guardar y generar reportes a partir de varios elementos.

Los structs se utilizan para representar entidades del sistema, como Libro y Categoria. Cada struct agrupa atributos relacionados dentro de una misma estructura, permitiendo trabajar los datos de forma más organizada.

---

## 7. Encapsulación

La encapsulación se implementó mediante atributos privados en los structs. En Go, los atributos que inician con letra minúscula solo pueden ser accedidos dentro del mismo paquete.

Por ejemplo, en el struct Libro se utilizan atributos como `titulo`, `autor`, `categoria`, `anio` y `disponible`. Estos atributos no se modifican directamente desde otros paquetes, sino mediante métodos getter y setter.

Esto permite proteger la información interna del sistema y controlar la forma en que los datos son consultados o modificados.

---

## 8. Métodos getter y setter

Los métodos getter permiten consultar los valores privados de una estructura. En el caso del libro, se implementaron métodos para obtener el ID, título, autor, categoría, año y disponibilidad.

Los métodos setter permiten modificar valores de forma controlada. Además, estos métodos incluyen validaciones para evitar datos incorrectos, como títulos vacíos, autores vacíos, categorías vacías o años inválidos.

Esto mejora la seguridad y consistencia de los datos dentro del sistema.

---

## 9. Manejo de errores

El sistema utiliza manejo de errores mediante el tipo `error` y la función `errors.New`. Esto permite detectar situaciones incorrectas durante la ejecución del programa.

Por ejemplo, si se intenta crear un libro con un título vacío o un año inválido, el sistema devuelve un error en lugar de aceptar información incorrecta.

El manejo de errores ayuda a que el programa sea más robusto, controlado y fácil de depurar.

---

## 10. Uso de interfaces

En el paquete de búsquedas se implementó una interfaz llamada `LibroBuscable`. Esta interfaz define los métodos que debe cumplir cualquier elemento que pueda ser buscado dentro del sistema.

El uso de interfaces permite desacoplar el módulo de búsquedas del struct Libro. De esta manera, las funciones de búsqueda pueden trabajar con cualquier tipo que implemente los métodos requeridos.

Esto representa una aplicación del polimorfismo mediante interfaces en Go y permite construir un código más flexible y reutilizable.

---

## 11. Persistencia en JSON

El sistema permite guardar los libros en un archivo JSON ubicado en la carpeta `data`. Esta funcionalidad permite conservar la información después de la ejecución del programa.

También se implementó la carga de datos desde el archivo JSON, demostrando que el sistema puede recuperar información previamente almacenada.

---

## 12. Reportes del sistema

El sistema genera reportes básicos del catálogo. El reporte general muestra:

* Total de libros registrados.
* Cantidad de libros disponibles.
* Cantidad de libros no disponibles.

Además, el reporte por categoría permite conocer cuántos libros pertenecen a una categoría específica.

---

## 13. Evidencia de ejecución

El sistema se ejecuta con el siguiente comando:

go run ./cmd

Durante la ejecución se evidencia:

* Registro correcto de una categoría.
* Registro correcto de un libro.
* Cambio de disponibilidad del libro.
* Búsqueda por título.
* Búsqueda por categoría.
* Guardado de libros en JSON.
* Carga de libros desde JSON.
* Generación de reporte general.
* Generación de reporte por categoría.

---

## 14. Relación con la rúbrica

El proyecto cumple con el Paso 1 porque continúa la planificación realizada en el Aprendizaje Autónomo 1 y presenta un avance significativo del desarrollo del sistema.

También cumple con el Paso 2 porque implementa funcionalidades relacionadas con encapsulación, manejo de errores e interfaces. Además, el código contiene comentarios en las partes importantes para facilitar su comprensión.

La implementación es clara, objetiva y fundamentada, ya que cada paquete cumple una responsabilidad específica dentro del sistema.

---

## 15. Conclusión

En conclusión, el avance desarrollado demuestra la transición de la planificación hacia una implementación funcional del Sistema de Gestión de Libros Electrónicos. El proyecto aplica estructuras de datos, structs, slices, métodos, encapsulación, manejo de errores, interfaces, persistencia en JSON y generación de reportes.

El sistema se encuentra organizado en paquetes, documentado en el README y cargado en GitHub para su revisión. Aunque todavía puede ampliarse en futuras etapas, este avance cumple con los requisitos principales del Aprendizaje Autónomo 2 y demuestra la aplicación práctica de los contenidos estudiados en clase.
