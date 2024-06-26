// 5. SGPC - Sistema de gestión de proyectos creativos

// Imagina que eres parte de un equipo de desarrollo de software en una empresa de tecnología creativa. Se te ha asignado la tarea de crear un sistema para gestionar proyectos creativos.
// El sistema debe permitir la creación de nuevos proyectos, asignación de miembros del equipo a proyectos, seguimiento del progreso del proyecto y generación de informes de estado del proyecto. Utiliza tu imaginación para definir cómo serán los proyectos creativos y cómo se gestionarán en este sistema.

// Requisitos:

// Implementar un sistema para la creación de nuevos proyectos creativos.
// Permitir la asignación de miembros del equipo a proyectos.
// Seguimiento del progreso del proyecto, incluyendo hitos y fechas límite.
// Generación de informes de estado del proyecto, que muestren el progreso, los miembros del equipo asignados y cualquier problema o desafío que surja durante el desarrollo del proyecto.
// Utiliza tu creatividad para definir otros requisitos y funcionalidades que consideres importantes para la gestión eficaz de proyectos creativos.

package main

import "fmt"

type Progreso struct {
	Estado      string
	Hitos       []string
	FechaLimite string
}

type Proyecto struct {
	Nombre      string
	Descripcion string
	Responsable string
	Miembros    []string
	Progreso    Progreso
}

type Proyectos struct {
	proyectos []Proyecto
}

// el progreso se inicializa en estado "pendiente", sin miembros asignados, y como primer hito se pone la creacion del proyecto con la fecha de la misma.
//se asigna como fecha limita, la fecha de entrega del proyecto
func (p *Proyectos) agregarProyecto(nombre, descripcion, responsable, fechaCreacion, fechaEntrega string) {
	nuevoProyecto := Proyecto{
		Nombre:      nombre,
		Descripcion: descripcion,
		Responsable: responsable,
		Miembros:    []string{},
		Progreso: Progreso{
			Estado:      "Pendiente",
			Hitos:       []string{"Proyecto creado." + fechaCreacion},
			FechaLimite: fechaEntrega,
		},
	}
	p.proyectos = append(p.proyectos, nuevoProyecto)
	fmt.Println("***")
	fmt.Printf("Nuevo proyecto agregado exitosamente: %s.\n", nombre)
	fmt.Println("***")

}

//buca por el nombre dentro del arreglo proyectos, devolviendo un punturo del proyecto si lo encuentra o nil en caso contrario.
//también devuelve un string para verificar si se encontró
func buscarProyecto(nombre string, p *Proyectos) (*Proyecto, string) {
	for i := range p.proyectos {
		if p.proyectos[i].Nombre == nombre {
			return &p.proyectos[i], ""
		}
	}
	return nil, "Proyecto no encontrado"
}

// se asigna un miembro al proyecto. Si es el primer miembro en asignarse al proyecto, el estado del proyecto pasa a "en progreso" y se define el hito de "Inicio de proyecto". 
// después de la primera asignación, se añaden miembros al proyecto creando el hito de creación con la fecha de asignación del nuevo miembro
func asignarMiembro(nombre string, miembro string, fecha string, p *Proyectos) {
	proyecto, error := buscarProyecto(nombre, p)
	fmt.Println("***")
	if proyecto != nil {
		proyecto.Miembros = append(proyecto.Miembros, miembro)
		proyecto.Progreso.Hitos = append(proyecto.Progreso.Hitos, "Se asigna al equipo: "+miembro+"."+fecha)
		fmt.Println()
		fmt.Println(proyecto.Nombre)
		fmt.Printf(" -%s añadido como miembro exitosamente. %s", miembro, fecha)
		fmt.Println()
		if len(proyecto.Miembros) == 1 {
			proyecto.Progreso.Estado = "En progreso"
			proyecto.Progreso.Hitos = append(proyecto.Progreso.Hitos, "Inicio de Proyecto."+fecha)
			fmt.Println(" -Felicidades!!Tu primer miembro en el equipo, oficialmente el proyecto comenzó")
		}
		return
	} else {
		fmt.Println(error)
	}
	fmt.Println("***")
}

// Da un informe de la situación del proyecto incluyendo estado, hitos, fecha Limite, 
func seguimientoProgreso(nombre string, p *Proyectos) {
	fmt.Println("***")
	proyecto, encontrado := buscarProyecto(nombre, p)
	if proyecto != nil {
		fmt.Printf("Progreso del proyecto %s: \n", proyecto.Nombre)
		fmt.Printf("-Estado: %s\n", proyecto.Progreso.Estado)
		fmt.Println("-Hitos: ")
		for i := range proyecto.Progreso.Hitos {
			fmt.Printf(" *Hito %d: %s\n", i+1, proyecto.Progreso.Hitos[i])
		}
		fmt.Printf("-Fecha Límite")
		fmt.Printf(" *Fecha Entrega: %s\n", proyecto.Progreso.FechaLimite)

	} else {
		fmt.Println(encontrado)
	}
	fmt.Println("***")
}

func mostrarProyectos(p *Proyectos) {
	fmt.Println("***")
	fmt.Println("Lista Proyectos:")
	if len(p.proyectos) == 0 {
		fmt.Println(" Aún no tienes proyectos.")
	} else {
		for i := range p.proyectos {
			fmt.Printf(" %d.%s\n", i+1, p.proyectos[i].Nombre)
		}
		fmt.Println("***")
	}
}

func main() {
	proyectos := Proyectos{}
	mostrarProyectos(&proyectos)
	proyectos.agregarProyecto("Proyecto Alpha", "Descripción del Proyecto Alpha", "Responsable A", "2024-06-03", "2024-12-31")
	seguimientoProgreso("Proyecto Alpha", &proyectos)
	proyectos.agregarProyecto("Proyecto Betha", "Descripción del Proyecto Betha", "Responsable B", "2023-06-09", "2023-12-31")
	asignarMiembro("Proyecto Alpha", "Miembro 1", "2024-06-11", &proyectos)
	seguimientoProgreso("Proyecto Alpha", &proyectos)
	asignarMiembro("Proyecto Alpha", "Miembro 2", "2024-06-12", &proyectos)
	seguimientoProgreso("Proyecto Alpha", &proyectos)
	asignarMiembro("Proyecto Betha", "Miembro 1", "2023-08-15", &proyectos)
	asignarMiembro("Proyecto Betha", "Miembro 2", "2023-09-24", &proyectos)
	seguimientoProgreso("Proyecto Betha", &proyectos)
	mostrarProyectos(&proyectos)
}
