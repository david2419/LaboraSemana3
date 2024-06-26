package main

import "testing"

// func crearTarea(nombre, descripcion, responsable string) *Tarea {
// 	nuevaTarea := &Tarea{
// 		nombre:      nombre,
// 		descripcion: descripcion,
// 		responsable: responsable,
// 		estado:      "pendiente",
// 	}
// 	return nuevaTarea

// }

// func actualizarEstado(tarea *Tarea, estado string) {
// 	tarea.estado = estado
// }

func TestCrearTarea(t *testing.T) {
	nuevaTarea := crearTarea("NuevaTarea", "Descripcion", "Responsable")
	tareas := []*Tarea{nuevaTarea}
	if tareas[0].nombre != "NuevaTarea" || len(tareas) != 1 {
		t.Errorf("CrearTarea falló, se esperaba que NuevaTarea tuviera nombre: NuevaTarea y que tareas tuviera longitud 1. En su lugar se obtuvo %s y %v", nuevaTarea.nombre, len(tareas))
	}
}

func TestActualizarEstado(t *testing.T) {
	nuevaTarea := crearTarea("NuevaTarea", "Descripcion", "Responsable")
	actualizarEstado(nuevaTarea, "En Proceso")
	if nuevaTarea.estado != "En Proceso" {
		t.Errorf("ActualizarEstado falló, se esperaba que el estado de nuevaTarea fuera En Proceso, en su lugar resulta: %s", nuevaTarea.estado)
	}
}
