Proyecto Final - Implementación de la Unidad 4
Sistema de Gestión de Libros Electrónicos

Estudiante: Matthew Adrian Narváez Paredes
Asignatura: Programación con Go/Golang
Actividad: Evaluación en Contacto con el Docente - Proyecto Final
Fecha: Junio de 2026
Repositorio: https://github.com/manarvaez207/gestion-libros-electronicos

1. Introducción

El presente informe documenta la implementación final del Sistema de Gestión de Libros Electrónicos, desarrollado en el lenguaje de programación Go/Golang. El proyecto inició como una propuesta de planificación de software, posteriormente se convirtió en una aplicación funcional por consola y finalmente evolucionó hacia una aplicación web basada en servicios HTTP y comunicación mediante JSON.

Esta etapa corresponde a la Implementación de la Unidad 4, en la cual se integran servicios web, serialización JSON, pruebas de software y conceptos de concurrencia en la gestión de solicitudes.

El sistema permite registrar, consultar, actualizar, eliminar, buscar y reportar libros electrónicos, utilizando una arquitectura modular organizada en paquetes internos.

2. Selección de la aplicación

La aplicación seleccionada fue un Sistema de Gestión de Libros Electrónicos. Esta elección se justifica porque permite resolver una necesidad real relacionada con la organización y consulta de recursos educativos digitales.

El sistema es factible para el alcance académico porque permite aplicar los contenidos estudiados en la materia, como:

Funciones.
Paquetes y módulos.
Structs.
Métodos.
Encapsulación.
Getters y setters.
Manejo de errores.
Interfaces.
Servicios web.
JSON.
Pruebas de software.
Concurrencia básica.

Además, se conecta con el tema del impacto de las nuevas tecnologías en la sociedad, ya que representa una base para bibliotecas digitales más accesibles, automatizadas e interconectadas.

3. Objetivo de la aplicación

Desarrollar una aplicación web en Go para la gestión de libros electrónicos, que permita administrar información bibliográfica mediante servicios web, aplicando programación modular, programación orientada a objetos en Go, manejo de errores, interfaces, persistencia JSON y pruebas de software.

4. Diseño y desarrollo del sistema

El sistema fue diseñado con una estructura modular para separar responsabilidades. Cada paquete cumple una función específica dentro de la aplicación.

gestion-libros-electronicos/
│
├── cmd/
│   └── main.go
│
├── data/
│   └── libros.json
│
├── internal/
│   ├── libros/
│   │   ├── libro.go
│   │   └── libro_test.go
│   │
│   ├── categorias/
│   │   ├── categoria.go
│   │   └── categoria_test.go
│   │
│   ├── busquedas/
│   │   ├── busquedas.go
│   │   └── busquedas_test.go
│   │
│   ├── persistencia/
│   │   └── persistencia.go
│   │
│   ├── reportes/
│   │   └── reportes.go
│   │
│   └── web/
│       ├── handlers.go
│       └── handlers_test.go
│
├── README.md
└── go.mod

5. Diagrama general de clases y módulos

classDiagram
    class Libro {
        -int id
        -string titulo
        -string autor
        -string categoria
        -int anio
        -bool disponible
        +ID() int
        +Titulo() string
        +Autor() string
        +Categoria() string
        +Anio() int
        +Disponible() bool
        +SetTitulo(titulo string) error
        +SetAutor(autor string) error
        +SetCategoria(categoria string) error
        +SetAnio(anio int) error
        +CambiarDisponibilidad(disponible bool)
        +Descripcion() string
    }

    class Categoria {
        -int id
        -string nombre
        +ID() int
        +Nombre() string
        +SetNombre(nombre string) error
    }

    class LibroBuscable {
        <<interface>>
        +Titulo() string
        +Autor() string
        +Categoria() string
        +Descripcion() string
    }

    class App {
        -[]Libro libros
        -[]Categoria categorias
        -string rutaDatos
        -sync.Mutex mu
        +Rutas() http.Handler
    }

    Libro ..|> LibroBuscable
    App --> Libro
    App --> Categoria

    6. Servicios web implementados

La aplicación web fue implementada usando el paquete estándar net/http. La API se ejecuta localmente en:

http://localhost:8080

Los servicios web implementados son:

| Método | Ruta                                      | Funcionalidad                          |
| ------ | ----------------------------------------- | -------------------------------------- |
| GET    | `/`                                       | Muestra información general de la API. |
| GET    | `/libros`                                 | Lista todos los libros registrados.    |
| GET    | `/libros/{id}`                            | Consulta un libro por ID.              |
| POST   | `/libros`                                 | Registra un nuevo libro.               |
| PUT    | `/libros/{id}`                            | Actualiza un libro existente.          |
| DELETE | `/libros/{id}`                            | Elimina un libro por ID.               |
| PATCH  | `/libros/{id}/disponibilidad`             | Cambia la disponibilidad de un libro.  |
| GET    | `/libros/buscar?titulo=go`                | Busca libros por título.               |
| GET    | `/categorias`                             | Lista las categorías registradas.      |
| POST   | `/categorias`                             | Registra una nueva categoría.          |
| GET    | `/reportes/general`                       | Genera un reporte general.             |
| GET    | `/reportes/categoria?nombre=Programacion` | Genera un reporte por categoría.       |

Con estos servicios se cumple el requisito de implementar al menos 8 servicios web de distintas funcionalidades del aplicativo.

7. Uso de JSON

El sistema utiliza JSON para recibir y responder información en los servicios web. Esto permite que la aplicación pueda comunicarse con otros sistemas, navegadores, clientes HTTP o futuras interfaces gráficas.

Ejemplo de respuesta JSON del endpoint /libros:

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

8. Concurrencia en la gestión de solicitudes

Go permite manejar solicitudes HTTP de forma concurrente. Cada solicitud recibida por el servidor puede ser atendida de manera independiente.

Para evitar problemas al modificar la lista de libros y categorías, se utilizó sync.Mutex dentro de la estructura App. Esto permite bloquear temporalmente el acceso a los datos compartidos cuando una operación los está modificando.

Además, se implementó una prueba de concurrencia en internal/web/handlers_test.go, donde se simulan varias solicitudes simultáneas al endpoint /libros usando goroutines y sync.WaitGroup.

Esto demuestra que el sistema considera la concurrencia en la gestión de solicitudes, cumpliendo con lo solicitado en la Unidad 4.

9. Pruebas de software realizadas

Se implementaron pruebas unitarias, pruebas de integración y una prueba básica de concurrencia.

9.1 Pruebas unitarias

Las pruebas unitarias verifican funciones y métodos específicos del sistema.

Archivos implementados:

internal/libros/libro_test.go
internal/categorias/categoria_test.go
internal/busquedas/busquedas_test.go

Estas pruebas validan:

Creación correcta de libros.
Validación de errores en datos inválidos.
Cambio de disponibilidad.
Validación de año inválido.
Creación correcta de categorías.
Cambio de nombre de categoría.
Búsqueda por título.
Búsqueda por categoría.
Uso de interfaces en las búsquedas.
9.2 Pruebas de integración

Las pruebas de integración se implementaron en:

internal/web/handlers_test.go

Estas pruebas verifican que los servicios web funcionen correctamente mediante solicitudes HTTP simuladas con httptest.

Se probaron los siguientes servicios:

GET /libros
POST /libros
PUT /libros/1
PATCH /libros/1/disponibilidad
DELETE /libros/1
GET /reportes/general
9.3 Prueba de concurrencia

También se implementó una prueba que ejecuta múltiples solicitudes simultáneas al endpoint /libros.

Esta prueba permite comprobar que el sistema puede responder a varias solicitudes concurrentes sin fallar.

10. Resultados de las pruebas

Para ejecutar las pruebas se utilizaron los comandos:

go fmt ./...
go test ./...

El resultado obtenido fue satisfactorio:

ok  sistema-libros/internal/busquedas
ok  sistema-libros/internal/categorias
ok  sistema-libros/internal/libros
ok  sistema-libros/internal/web

Esto confirma que las pruebas unitarias, pruebas de integración y prueba de concurrencia fueron ejecutadas correctamente.

11. Pruebas de aceptación manuales

Además de las pruebas automatizadas, se realizaron pruebas manuales desde el navegador y desde PowerShell.

Se verificaron los siguientes endpoints:

http://localhost:8080
http://localhost:8080/libros
http://localhost:8080/categorias
http://localhost:8080/reportes/general
http://localhost:8080/reportes/categoria?nombre=Programación
http://localhost:8080/libros/buscar?titulo=go
http://localhost:8080/libros/1

También se probaron operaciones con PowerShell:

Crear libro con POST.
Actualizar libro con PUT.
Cambiar disponibilidad con PATCH.
Eliminar libro con DELETE.

Estas pruebas permitieron comprobar que el sistema funciona correctamente desde el punto de vista del usuario final.

12. Errores encontrados y solución aplicada

Durante el desarrollo se encontraron algunos errores técnicos:

Puerto 8080 ocupado

Al ejecutar el servidor, apareció el error:

listen tcp :8080: bind: Solo se permite un uso de cada dirección de socket

Este error ocurrió porque una ejecución anterior del servidor seguía usando el puerto 8080.

La solución fue identificar el proceso con:

netstat -ano | findstr :8080

Y finalizarlo con:

taskkill /PID <PID> /F

Problemas de codificación en PowerShell

En algunas respuestas de PowerShell, las palabras con tilde podían mostrarse con caracteres extraños. Esto no afectó la lógica del sistema, pero para las pruebas se recomienda usar datos sin tildes cuando se ejecuten comandos desde la terminal.

13. Optimización realizada

La principal optimización fue organizar el sistema por módulos y separar la lógica de negocio de la capa web.

Además:

Se usó sync.Mutex para proteger datos compartidos.
Se usó JSON como formato estándar de intercambio.
Se implementó validación de datos.
Se agregaron pruebas automatizadas.
Se mantuvo el código con comentarios en funciones relevantes.
Se actualizó el repositorio de GitHub con todos los cambios.
14. Limitaciones del sistema

El sistema cumple con los objetivos académicos, pero presenta algunas limitaciones:

La persistencia se realiza mediante archivo JSON local y no mediante base de datos.
La API funciona localmente en localhost:8080.
No existe autenticación de usuarios.
No hay roles de administrador, docente o estudiante.
No posee interfaz gráfica web.
No está desplegado en la nube.
No maneja préstamos digitales avanzados.

Estas limitaciones representan oportunidades de mejora para una versión futura.

15. Conclusión

En conclusión, el proyecto permitió integrar los conocimientos desarrollados durante la materia mediante una aplicación funcional en Go. El sistema evolucionó desde una aplicación de consola hacia una API Web con servicios HTTP, respuestas JSON, pruebas de software y control básico de concurrencia.

La implementación de servicios web permitió demostrar cómo un sistema de gestión de libros electrónicos puede convertirse en una solución más flexible, accesible e interoperable. Además, el uso de pruebas unitarias, integración y concurrencia permitió mejorar la calidad del código y documentar el funcionamiento del sistema.

Este proyecto también permite visualizar un futuro donde las bibliotecas digitales sean más inteligentes, automatizadas e interconectadas, contribuyendo al acceso a la información y al fortalecimiento de la educación digital.
