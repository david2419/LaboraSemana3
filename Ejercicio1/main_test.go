package main

import "testing"

func TestCrearProducto(t *testing.T) {
	prod := CrearProducto("Producto1", 100, 10)
	if prod.nombre != "Producto1" || prod.precio != 100 || prod.disponible != 10 {
		t.Errorf("CrearProducto falló, se esperaba nombre: Producto1, precio: 100, disponible: 10, pero se obtuvo nombre: %s, precio: %d, disponible: %d", prod.nombre, prod.precio, prod.disponible)
	}
}

func TestAgregarProducto(t *testing.T) {
	inventario := []Producto{}
	prod := CrearProducto("Producto1", 100, 10)
	inventario = AgregarProducto(inventario, prod)
	if len(inventario) != 1 || inventario[0].nombre != "Producto1" || inventario[0].precio != 100 || inventario[0].disponible != 10 {
		t.Errorf("AgregarProducto falló, se esperaba un inventario con un producto con nombre: Producto1, precio: 100, disponible: 10, pero se obtuvo: %+v", inventario)
	}
}

func TestActualizarDisponible(t *testing.T) {
	inventario := []Producto{
		{"Producto1", 100, 10},
	}
	actualizarDisponible(inventario, "Producto1", 5)
	if inventario[0].disponible != 5 {
		t.Errorf("actualizarDisponible falló, se esperaba disponible: 5, pero se obtuvo: %d", inventario[0].disponible)
	}
}

func TestEliminarProducto(t *testing.T) {
	inventario := []Producto{
		{"Producto1", 100, 10},
		{"Producto2", 200, 20},
	}
	inventario = eliminarProducto(inventario, "Producto1")
	if len(inventario) != 1 || inventario[0].nombre != "Producto2" {
		t.Errorf("eliminarProducto falló, se esperaba un inventario con un producto con nombre: Producto2, pero se obtuvo: %+v", inventario)
	}
}

