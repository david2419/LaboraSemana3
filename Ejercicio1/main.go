package main

import "fmt"

// GESTIÓN DE INVENTARIOS
// Una pequeña tienda necesita un programa para gestionar su inventario de productos.
//El programa debe permitir la adición de nuevos productos, la actualización de la cantidad disponible de un producto,
//la eliminación de productos y la visualización del inventario completo.

// Requisitos:

// Implementar una estructura de datos para almacenar la información de cada producto (nombre, precio y cantidad disponible).
// Permitir la adición de nuevos productos con sus respectivas cantidades.
// Permitir la actualización de la cantidad disponible de un producto existente.
// Permitir la eliminación de productos del inventario.
// Mostrar el inventario completo, incluyendo el nombre, precio y cantidad disponible de cada producto.

type Producto struct {
	nombre     string
	precio     int
	disponible int
}

func CrearProducto(nombr string, prec int, disp int) *Producto {
	nuevoProducto := &Producto{
		nombre:     nombr,
		precio:     prec,
		disponible: disp,
	}
	return nuevoProducto
}

func AgregarProducto(inventario []Producto, nuevoProducto *Producto) []Producto {
	inventario = append(inventario, *nuevoProducto)
	return inventario
}

func actualizarDisponible(inventario []Producto, productoBuscado string, nDisponible int) {
	for i := range inventario {
		if inventario[i].nombre == productoBuscado {
			inventario[i].disponible = nDisponible
			return
		}
	}
}

func eliminarProducto(inventario []Producto, productoBuscado string) []Producto {
	for i := range inventario {
		if inventario[i].nombre == productoBuscado {
			inventario = append(inventario[:i], inventario[i+1:]...)
			return inventario
		}
	}
	return inventario
}

func mostrarProductos(inventario []Producto) {
	fmt.Println("Inventario: ")
	for _, producto := range inventario {
		fmt.Printf("Producto: %s, Precio: %d, Disponible: %d\n", producto.nombre, producto.precio, producto.disponible)

	}
}

func main() {
	var inventario []Producto

	Pastel := CrearProducto("Pastel", 5500, 4)
	fmt.Println(Pastel.nombre)
	Oblea := CrearProducto("Oblea", 8000, 4)
	Jugo := CrearProducto("Jugo", 12000, 4)

	inventario = AgregarProducto(inventario, Pastel)
	inventario = AgregarProducto(inventario, Oblea)
	inventario = AgregarProducto(inventario, Jugo)

	mostrarProductos(inventario)

	actualizarDisponible(inventario, "Pastel", 12)
	fmt.Println("")
	mostrarProductos(inventario)

	inventario = eliminarProducto(inventario, "Oblea")

	fmt.Println("")
	mostrarProductos(inventario)

}
