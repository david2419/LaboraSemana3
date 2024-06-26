//ejercicio 3
// Una biblioteca necesita un sistema para gestionar su colección de libros. El programa debe permitir la adición de nuevos libros, la búsqueda de libros por título o autor, la actualización del estado de un libro (disponible o prestado) y la eliminación de libros.

// Requisitos:

// Implementar una estructura de datos para almacenar la información de cada libro (título, autor, género y estado).
// Permitir la adición de nuevos libros a la colección.
// Permitir la búsqueda de libros por título o autor.
// Permitir la actualización del estado de un libro a "disponible" o "prestado".
// Permitir la eliminación de libros de la colección.
package main

import "fmt"

type Libro struct {
	Titulo     string
	Autor      string
	Genero     string
	Disponible bool
}

func crearLibro(titulo, autor, genero string, disponible bool) Libro {
	nuevoLibro := Libro{
		Titulo:     titulo,
		Autor:      autor,
		Genero:     genero,
		Disponible: disponible,
	}
	return nuevoLibro
}

func agregarLibro(coleccion map[string]Libro, titulo, autor, genero string, disponible bool) {
	nuevoLibro := crearLibro(titulo, autor, genero, disponible)
	coleccion[nuevoLibro.Titulo] = nuevoLibro
}

func busquedaPorTitulo(titulo string, coleccion map[string]Libro) (Libro, bool) {
	libro, encontrado := coleccion[titulo]
	return libro, encontrado
}

func busquedaPorAutor(autor string, coleccion map[string]Libro) []Libro {
	libros := []Libro{}
	for _, libro := range coleccion {
		if libro.Autor == autor {
			libros = append(libros, libro)
		}
	}
	return libros
}

func actualizacionEstado(titulo string, disponible bool, coleccion map[string]Libro) {
	libro, encontrado := busquedaPorTitulo(titulo, coleccion)
	if encontrado {
		libro.Disponible = disponible
	}
}

func eliminarLibro(titulo string, coleccion map[string]Libro) {
	libro, encontrado := busquedaPorTitulo(titulo, coleccion)
	if encontrado {
		delete(coleccion, libro.Titulo)
	}
}

func mostrarColeccion(coleccion map[string]Libro) {
	fmt.Println("Colección de libros:")
	for _, libro := range coleccion {
		estado := "disponible"
		if !libro.Disponible {
			estado = "prestado"
		}
		fmt.Printf("- Titulo: %s, Autor: %s, Genero: %s, Estado: %s\n", libro.Titulo, libro.Autor, libro.Genero, estado)
	}
}

func main() {
	coleccion := make(map[string]Libro)

	// Agregar algunos libros a la colección
	agregarLibro(coleccion, "El principito", "Antoine de Saint-Exupéry", "Fábula", true)
	agregarLibro(coleccion, "Cien años de soledad", "Gabriel García Márquez", "Realismo mágico", true)
	agregarLibro(coleccion, "1984", "George Orwell", "Distopía", true)
	agregarLibro(coleccion, "Matar a un ruiseñor", "Harper Lee", "Ficción", true)
	agregarLibro(coleccion, "Orgullo y prejuicio", "Jane Austen", "Romance", true)

	// Mostrar la colección completa
	mostrarColeccion(coleccion)

	// Actualizar el estado de un libro
	actualizacionEstado("El principito", false, coleccion)

	// Mostrar la colección después de actualizar el estado de un libro
	mostrarColeccion(coleccion)

	// Eliminar un libro de la colección
	eliminarLibro("1984", coleccion)

	// Mostrar la colección después de eliminar un libro
	mostrarColeccion(coleccion)
}
