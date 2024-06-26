// 4. Diccionario de contactos

// Un usuario necesita un programa para gestionar sus contactos. El programa debe permitir la adición de nuevos contactos, la búsqueda de contactos por nombre o número de teléfono, la actualización de la información de un contacto y la eliminación de contactos.

// Requisitos:

// Implementar una estructura de datos para almacenar la información de cada contacto (nombre, número de teléfono, correo electrónico y dirección).
// Permitir la adición de nuevos contactos con su información completa.
// Permitir la búsqueda de contactos por nombre o número de teléfono.
// Permitir la actualización de la información de un contacto (teléfono, correo electrónico o dirección).
// Permitir la eliminación de contactos de la lista.

package main

import "fmt"

type Contacto struct {
	Nombre    string
	Telefono  string
	Correo    string
	Direccion string
}

func agregarContacto(agenda map[string]Contacto, nombre, telefono, correo, direccion string) {
	nuevoContacto := Contacto{
		Nombre:    nombre,
		Telefono:  telefono,
		Correo:    correo,
		Direccion: direccion,
	}

	agenda[nombre] = nuevoContacto
	fmt.Printf("%s agregado a la agenda. \n", nombre)
}

func busquedaContactoNombre(agenda map[string]Contacto, nombre string) {
	fmt.Println("Buscar contacto por nombre : ")
	contacto, encontrado := agenda[nombre]
	if !encontrado {
		fmt.Printf("%s no está en tus contactos \n", nombre)
	}
	fmt.Println(contacto, encontrado)

}

func busquedaContactoTelefono(agenda map[string]Contacto, telefono string) {
	fmt.Println("Buscar contacto por telefono : ")
	encontrado := false
	var contactoEncontrado Contacto
	for _, contacto := range agenda {
		if contacto.Telefono == telefono {
			encontrado = true
			contactoEncontrado = contacto
		}
	}
	fmt.Println(contactoEncontrado, encontrado)
}

func actualizarContacto(agenda map[string]Contacto, nombre, telefono, correo, direccion string) {
	contactoActualizacion, encontrado := agenda[nombre]
	if encontrado {
		if telefono != "" {
			contactoActualizacion.Telefono = telefono
		}
		if correo != "" {
			contactoActualizacion.Correo = correo
		}
		if direccion != "" {
			contactoActualizacion.Direccion = direccion
		}
		agenda[nombre] = contactoActualizacion
		fmt.Printf("Contacto de %s ha sido actualizado.\n", nombre)
	} else {
		fmt.Println("Imposinler actualizar.")
	}
}

func eliminarContacto(agenda map[string]Contacto, nombre string) {

	_, encontrado := agenda[nombre]
	if encontrado {
		delete(agenda, nombre)
		fmt.Println(nombre, "eliminado de tu agenda.")
	}
}

func mostrarAgenda(agenda map[string]Contacto) {
	fmt.Println("AGENDA:")
	for contacto := range agenda {
		fmt.Println(agenda[contacto])
	}
}

func main() {
	agenda := make(map[string]Contacto)

	agregarContacto(agenda, "Juan Perez", "123456789", "juan@example.com", "Calle Falsa 123")
	agregarContacto(agenda, "Maria Gomez", "987654321", "maria@example.com", "Avenida Siempre Viva 456")

	mostrarAgenda(agenda)

	busquedaContactoNombre(agenda, "Juan Perez")
	busquedaContactoNombre(agenda, "asdf")
	busquedaContactoTelefono(agenda, "987654321")
	busquedaContactoTelefono(agenda, "2342342")

	actualizarContacto(agenda, "pepito", "", "", "")
	actualizarContacto(agenda, "Juan Perez", "", "", "CedroBolivar")

	mostrarAgenda(agenda)

	eliminarContacto(agenda, "JOSE SUAREZ")
	eliminarContacto(agenda, "Maria Gomez")

	mostrarAgenda(agenda)
}
