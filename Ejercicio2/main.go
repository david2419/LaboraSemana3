package main

import "fmt"

// 2. Sistema de gestión de tareas

// Un equipo de desarrollo necesita un programa para gestionar sus tareas pendientes. El programa debe permitir la creación de nuevas tareas, la asignación de responsables, la actualización del estado de una tarea y la visualización de todas las tareas pendientes.

// Requisitos:

// Implementar una estructura de datos para almacenar la información de cada tarea (nombre, descripción, responsable y estado).
// Permitir la creación de nuevas tareas con un estado inicial de "pendiente".
// Permitir la asignación de un responsable a cada tarea.
// Permitir la actualización del estado de una tarea a "en progreso" o "completada".
// Mostrar todas las tareas pendientes, incluyendo su nombre, descripción, responsable y estado.

type Tarea struct {
	nombre      string
	descripcion string
	responsable string
	estado      string
}

func crearTarea(nombre, descripcion, responsable string) *Tarea {
	nuevaTarea := &Tarea{
		nombre:      nombre,
		descripcion: descripcion,
		responsable: responsable,
		estado:      "pendiente",
	}
	return nuevaTarea

}

func actualizarEstado(tarea *Tarea, estado string) {
	tarea.estado = estado
}

func tareasPendientes(tareas []*Tarea) {
	for _, tarea := range tareas {
		if tarea.estado == "pendiente" {
			fmt.Printf("Nombre: %s, Descripcion: %s, Estado: %s\n", tarea.nombre, tarea.descripcion, tarea.estado)
		}
	}
}

func main() {

	tarea1 := crearTarea("Tarea 1", "Descripción tarea 1", "Joaquin")
	tarea2 := crearTarea("Tarea 2", "Descripción tarea 2", "David")
	tarea3 := crearTarea("Tarea 3", "Descripción tarea 3", "Ana")
	tareas := []*Tarea{tarea1, tarea2, tarea3}
	fmt.Println(tarea1)
	actualizarEstado(tarea1, "Terminado")
	fmt.Println(tarea1)
	fmt.Println(tarea2)
	fmt.Println("Tareas Pendientes---------------------------")
	tareasPendientes(tareas)

}
