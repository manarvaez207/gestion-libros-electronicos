# Sistema de Gestión de Libros Electrónicos

## Proyecto Final: El impacto de las nuevas tecnologías en la sociedad: visualización del futuro

**Asignatura:** Programación con Go/Golang
**Actividad:** Evaluación en Contacto con el Docente - Proyecto Final
**Estudiante:** Matthew Adrian Narváez Paredes
**Fecha:** Junio de 2026
**Repositorio:** https://github.com/manarvaez207/gestion-libros-electronicos

---

## 1. Introducción

Este proyecto corresponde al desarrollo final de un **Sistema de Gestión de Libros Electrónicos** implementado en el lenguaje de programación **Go/Golang**. El sistema inició como una propuesta de planificación en el Aprendizaje Autónomo 1, continuó como una aplicación funcional por consola en el Aprendizaje Autónomo 2 y finalmente evolucionó hacia una **API Web con servicios HTTP y respuestas en formato JSON**.

El proyecto se relaciona con el tema **“El impacto de las nuevas tecnologías en la sociedad: visualización del futuro”**, debido a que representa cómo las bibliotecas digitales, los servicios web y la gestión automatizada de información pueden transformar el acceso al conocimiento en contextos educativos, institucionales y sociales.

La idea principal del proyecto es que el desarrollo de sistemas web para la gestión de libros electrónicos permite visualizar un futuro donde las bibliotecas digitales sean más accesibles, inteligentes, automatizadas e interconectadas.

---

## 2. Objetivo del proyecto

Desarrollar un sistema de gestión de libros electrónicos en Go que permita registrar, consultar, actualizar, eliminar, buscar y reportar libros mediante una arquitectura modular, aplicando estructuras de datos, encapsulación, manejo de errores, interfaces, persistencia en JSON y servicios web.

---

## 3. Criterio de selección del tema

Se seleccionó el sistema de gestión de libros electrónicos porque es un tema creativo, enfocado y factible. Además, permite relacionar la programación con una problemática real: la organización y acceso eficiente a recursos educativos digitales.

El tema no se limita únicamente al almacenamiento de libros, sino que proyecta una visión futura basada en:

* Bibliotecas digitales inteligentes.
* Acceso remoto a recursos educativos.
* Automatización de catálogos.
* Servicios web para interoperabilidad.
* Uso de datos estructurados mediante JSON.
* Posible integración futura con inteligencia artificial y computación en la nube.

---

## 4. Alcance del proyecto por semanas

| Semana | Unidad   | Tema relacionado | Alcance desarrollado                                                                  |
| ------ | -------- | ---------------- | ------------------------------------------------------------------------------------- |
| S1     | Unidad 1 | Tema 1           | Selección del sistema de gestión y definición del problema.                           |
| S2     | Unidad 1 | Tema 2           | Planificación del sistema, objetivo, módulos y estructura inicial.                    |
| S3     | Unidad 2 | Tema 3           | Uso de estructuras de datos como slices para manejar listas de libros.                |
| S4     | Unidad 2 | Tema 4           | Implementación de structs, métodos y funciones constructoras.                         |
| S5     | Unidad 3 | Tema 5           | Aplicación de encapsulación, métodos getter, métodos setter y manejo de errores.      |
| S6     | Unidad 3 | Tema 6           | Implementación de interfaces y polimorfismo para búsquedas.                           |
| S7     | Unidad 4 | Tema 7           | Generación de servicios web con `net/http` y serialización JSON.                      |
| S8     | Unidad 4 | Tema 8           | Integración final, pruebas, documentación, GitHub, presentación y video demostrativo. |

---

## 5. Tecnologías utilizadas

* **Go/Golang:** lenguaje principal del proyecto.
* **Visual Studio Code:** editor de código.
* **Git:** control de versiones.
* **GitHub:** repositorio remoto del proyecto.
* **JSON:** formato de almacenamiento e intercambio de datos.
* **net/http:** paquete estándar de Go para crear servicios web.
* **PowerShell:** ejecución de comandos y pruebas de servicios.

---

## 6. Estructura del proyecto

```text
gestion-libros-electronicos/
│
├── cmd/
│   └── main.go
│
├── data/
│   └── libros.json
│
├── docs/
│   └── avance_autonomo_2.md
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
│   ├── reportes/
│   │   └── reportes.go
│   │
│   └── web/
│       └── handlers.go
│
├── go.mod
└── README.md
```

---

## 7. Descripción de los módulos

### 7.1 Módulo `cmd`

Contiene el archivo `main.go`, que es el punto de entrada del sistema. En la versión final, este archivo inicia el servidor web en el puerto `8080`.

### 7.2 Módulo `libros`

Contiene la estructura principal `Libro`. Sus atributos están encapsulados mediante campos privados y se accede a ellos mediante métodos getter y setter.

Este módulo permite:

* Crear libros.
* Validar título, autor, categoría y año.
* Consultar datos mediante getters.
* Modificar datos mediante setters.
* Cambiar disponibilidad.
* Generar una descripción textual del libro.

### 7.3 Módulo `categorias`

Permite crear y validar categorías para clasificar los libros electrónicos. También aplica encapsulación mediante atributos privados y métodos de acceso.

### 7.4 Módulo `busquedas`

Implementa búsquedas por título y por categoría. Además, utiliza una interfaz llamada `LibroBuscable`, lo que permite aplicar polimorfismo y desacoplar el módulo de búsqueda de una estructura concreta.

### 7.5 Módulo `persistencia`

Permite guardar y cargar libros desde un archivo JSON. Esto permite conservar los datos entre ejecuciones del programa.

### 7.6 Módulo `reportes`

Genera reportes básicos del sistema, como:

* Total de libros.
* Libros disponibles.
* Libros no disponibles.
* Total de libros por categoría.

### 7.7 Módulo `web`

Contiene los servicios web del sistema. Este módulo permite que el aplicativo funcione como una API Web mediante el paquete estándar `net/http`.

---

## 8. Funcionalidades implementadas

El sistema final permite:

* Registrar libros electrónicos.
* Consultar todos los libros.
* Consultar un libro por ID.
* Actualizar la información de un libro.
* Eliminar un libro.
* Cambiar la disponibilidad de un libro.
* Registrar categorías.
* Consultar categorías.
* Buscar libros por título.
* Generar reporte general.
* Generar reporte por categoría.
* Guardar y cargar información usando JSON.
* Responder solicitudes mediante servicios web.

---

## 9. Servicios web implementados

La API se ejecuta en:

```text
http://localhost:8080
```

| Método | Ruta                                      | Descripción                                 |
| ------ | ----------------------------------------- | ------------------------------------------- |
| GET    | `/`                                       | Muestra información general de la API.      |
| GET    | `/libros`                                 | Lista todos los libros registrados.         |
| GET    | `/libros/{id}`                            | Consulta un libro por su ID.                |
| POST   | `/libros`                                 | Registra un nuevo libro.                    |
| PUT    | `/libros/{id}`                            | Actualiza completamente un libro.           |
| DELETE | `/libros/{id}`                            | Elimina un libro por ID.                    |
| PATCH  | `/libros/{id}/disponibilidad`             | Cambia la disponibilidad de un libro.       |
| GET    | `/libros/buscar?titulo=go`                | Busca libros por coincidencia en el título. |
| GET    | `/categorias`                             | Lista las categorías registradas.           |
| POST   | `/categorias`                             | Registra una nueva categoría.               |
| GET    | `/reportes/general`                       | Genera un reporte general del sistema.      |
| GET    | `/reportes/categoria?nombre=Programación` | Genera un reporte por categoría.            |

Con estos servicios se cumple el requisito de crear al menos 8 servicios web de distintas funcionalidades del aplicativo.

---

## 10. Ejemplos de respuestas JSON

### Listar libros

Ruta:

```text
GET http://localhost:8080/libros
```

Respuesta esperada:

```json
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
```

### Reporte general

Ruta:

```text
GET http://localhost:8080/reportes/general
```

Respuesta esperada:

```json
{
  "total_libros": 1,
  "libros_disponibles": 0,
  "libros_no_disponibles": 1
}
```

### Reporte por categoría

Ruta:

```text
GET http://localhost:8080/reportes/categoria?nombre=Programación
```

Respuesta esperada:

```json
{
  "categoria": "Programación",
  "total_libros": 1
}
```

---

## 11. Comandos de ejecución

Para formatear el código:

```powershell
go fmt ./...
```

Para verificar que el proyecto compile correctamente:

```powershell
go test ./...
```

Para ejecutar el servidor web:

```powershell
go run ./cmd
```

Al ejecutar el sistema debe mostrarse:

```text
Sistema de Gestión de Libros Electrónicos - API Web
Servidor iniciado en http://localhost:8080
```

Luego se puede acceder desde el navegador a:

```text
http://localhost:8080
http://localhost:8080/libros
http://localhost:8080/categorias
http://localhost:8080/reportes/general
```

---

## 12. Pruebas desde PowerShell

### Crear un libro con POST

```powershell
$json = '{"titulo":"Clean Code","autor":"Robert C. Martin","categoria":"Programacion","anio":2008,"disponible":true}'

Invoke-RestMethod -Uri "http://localhost:8080/libros" -Method Post -ContentType "application/json" -Body $json
```

### Actualizar un libro con PUT

```powershell
$json = '{"titulo":"Clean Code Actualizado","autor":"Robert C. Martin","categoria":"Programacion","anio":2009,"disponible":true}'

Invoke-RestMethod -Uri "http://localhost:8080/libros/2" -Method Put -ContentType "application/json" -Body $json
```

### Cambiar disponibilidad con PATCH

```powershell
$json = '{"disponible":false}'

Invoke-RestMethod -Uri "http://localhost:8080/libros/2/disponibilidad" -Method Patch -ContentType "application/json" -Body $json
```

### Eliminar un libro con DELETE

```powershell
Invoke-RestMethod -Uri "http://localhost:8080/libros/2" -Method Delete
```

---

## 13. Relación con las unidades de la asignatura

### Unidad 1: Fundamentos de Go

Se aplicaron funciones, paquetes, módulos y organización básica del código. También se estructuró el proyecto usando `go.mod` y carpetas separadas para mejorar la modularidad.

### Unidad 2: Estructuras de datos y objetos en Go

Se utilizaron estructuras como `struct` para representar libros y categorías. También se aplicaron slices para manejar listas de libros y categorías.

### Unidad 3: Programación orientada a objetos en Go

Se aplicó encapsulación mediante atributos privados, métodos getter para consultar información, métodos setter para modificar valores de forma controlada, manejo de errores mediante el tipo `error` e interfaces para desacoplar funcionalidades.

### Unidad 4: Servicios web, JSON e integración final

Se implementó una API Web usando `net/http`. Los datos se envían y reciben en formato JSON, permitiendo que el sistema pueda ser consultado desde un navegador o desde herramientas como PowerShell.

---

## 14. Análisis del software

El software evolucionó desde una planificación inicial hasta una API Web funcional. En la primera etapa se definió el sistema de gestión y sus módulos. En la segunda etapa se implementaron las estructuras principales, la lógica de negocio, las búsquedas, la persistencia y los reportes. En la etapa final se integró una capa web para exponer las funcionalidades mediante servicios HTTP.

El sistema demuestra varios patrones importantes:

* La modularidad permite separar responsabilidades.
* La encapsulación protege los datos internos.
* Los errores permiten controlar entradas inválidas.
* Las interfaces facilitan el desacoplamiento.
* JSON permite interoperabilidad con otros sistemas.
* Los servicios web permiten que la aplicación pueda ser consumida desde distintos clientes.

Esta evolución evidencia cómo una aplicación local puede transformarse en una solución más cercana a un sistema real usado en entornos educativos o institucionales.

---

## 15. Resultados obtenidos

Los principales resultados del proyecto son:

* Sistema de gestión de libros electrónicos funcional.
* API Web ejecutándose en `localhost:8080`.
* Más de 8 servicios web implementados.
* Serialización y deserialización mediante JSON.
* Persistencia de datos en `data/libros.json`.
* Registro, consulta, actualización y eliminación de libros.
* Búsqueda por título.
* Consulta de categorías.
* Reportes generales y por categoría.
* Código organizado en paquetes.
* Código subido y actualizado en GitHub.

Estos resultados demuestran que el sistema cumple el objetivo de gestionar información de libros electrónicos mediante una solución modular, documentada y conectada con servicios web.

---

## 16. Visualización del futuro

La visualización del futuro del proyecto plantea una biblioteca digital inteligente donde estudiantes, docentes e instituciones puedan acceder a libros electrónicos desde cualquier dispositivo. En este escenario, los servicios web permiten conectar aplicaciones móviles, plataformas educativas, repositorios digitales y sistemas administrativos.

A futuro, este sistema podría incorporar:

* Inteligencia artificial para recomendar libros.
* Búsquedas semánticas por tema o contenido.
* Autenticación de usuarios.
* Roles para administradores, docentes y estudiantes.
* Integración con bases de datos en la nube.
* Paneles de análisis de lectura.
* Préstamos digitales automatizados.
* Acceso desde aplicaciones móviles o plataformas web.

Esta visión muestra cómo las nuevas tecnologías pueden facilitar el acceso al conocimiento, reducir barreras educativas y mejorar la organización de recursos digitales.

---

## 17. Implicaciones del proyecto

El proyecto tiene implicaciones importantes en diferentes ámbitos:

### Ámbito educativo

Permite organizar libros electrónicos y facilitar su consulta, lo que mejora el acceso a recursos académicos.

### Ámbito tecnológico

Demuestra cómo una aplicación puede evolucionar hacia servicios web reutilizables e interoperables.

### Ámbito social

Contribuye a visualizar un futuro donde más personas puedan acceder al conocimiento mediante plataformas digitales.

### Ámbito institucional

Puede servir como base para sistemas de bibliotecas digitales en colegios, universidades o centros de formación.

---

## 18. Limitaciones del proyecto

Aunque el sistema cumple con los objetivos académicos, presenta algunas limitaciones:

* No utiliza una base de datos real, sino un archivo JSON local.
* No cuenta con autenticación de usuarios.
* No posee una interfaz gráfica avanzada.
* No está desplegado en la nube.
* No implementa roles de usuario.
* No tiene control de préstamos digitales por usuario.
* No incluye pruebas unitarias personalizadas.
* La API funciona localmente en `localhost:8080`.

Estas limitaciones no impiden el funcionamiento del sistema, pero representan oportunidades de mejora para versiones futuras.

---

## 19. Pros y contras

### Pros

* Código modular y organizado.
* Uso de Go como lenguaje eficiente.
* Implementación de servicios web.
* Manejo de JSON.
* Funcionalidades principales completas.
* Fácil ejecución local.
* Repositorio documentado en GitHub.

### Contras

* Persistencia limitada a archivo JSON.
* No cuenta con interfaz gráfica.
* No tiene autenticación.
* No está publicado en internet.
* No maneja múltiples usuarios simultáneos con base de datos.

---

## 20. Toma de decisiones

Durante el desarrollo se tomaron varias decisiones técnicas:

* Se utilizó Go por su claridad, rendimiento y soporte para servicios web.
* Se eligió `net/http` para evitar dependencias externas innecesarias.
* Se usó JSON porque es un formato estándar para servicios web.
* Se mantuvo una arquitectura modular para separar responsabilidades.
* Se conservaron los conceptos de encapsulación, interfaces y manejo de errores vistos en clase.
* Se decidió trabajar con archivo JSON local para mantener el proyecto simple, funcional y acorde al alcance académico.

Estas decisiones permiten que el sistema sea entendible, funcional y coherente con los contenidos estudiados.

---

## 21. Argumento principal

El desarrollo de sistemas web para la gestión de libros electrónicos demuestra cómo las nuevas tecnologías pueden transformar el acceso a la información, automatizar procesos educativos y proyectar un futuro con bibliotecas digitales más inteligentes, accesibles e interconectadas.

Este argumento se sustenta en que los servicios web permiten que la información no dependa únicamente de una aplicación local, sino que pueda ser consultada, actualizada y reutilizada por diferentes sistemas o dispositivos.

---

## 22. Conclusión

En conclusión, este proyecto permitió integrar los conocimientos adquiridos durante las unidades de la asignatura mediante el desarrollo de un Sistema de Gestión de Libros Electrónicos en Go. El sistema evolucionó desde una planificación inicial hasta una API Web funcional con servicios HTTP y datos en formato JSON.

Durante el desarrollo se aplicaron estructuras de datos, structs, métodos, encapsulación, getters, setters, manejo de errores, interfaces, persistencia y servicios web. Además, el proyecto permitió comprender cómo una solución tecnológica puede relacionarse con necesidades reales de acceso a la información y educación digital.

La principal dificultad fue integrar la lógica existente del sistema con una capa web funcional, manteniendo la organización modular y evitando errores en la serialización JSON. Sin embargo, esta dificultad permitió reforzar el aprendizaje sobre arquitectura de software, pruebas y control de versiones.

Finalmente, el proyecto demuestra que las nuevas tecnologías pueden contribuir a construir bibliotecas digitales más accesibles, automatizadas e inteligentes, abriendo oportunidades para mejorar el acceso al conocimiento en el futuro.
