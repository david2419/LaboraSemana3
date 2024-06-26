// 6. Juego de rol (RPG)

// Imagina que quieres crear un juego de rol (RPG) en Go. Tienes total libertad para diseñar y desarrollar este juego según tu imaginación. Puedes crear mundos, personajes, enemigos, misiones, habilidades y cualquier otra característica que desees incluir en el juego. Utiliza todos los conceptos que has aprendido sobre programación orientada a objetos en Go para desarrollar este emocionante proyecto.

// Requisitos:

// Diseñar e implementar la estructura del juego, incluyendo mundos, personajes, enemigos y misiones.
// Permitir que los jugadores controlen un personaje en el juego y puedan explorar el mundo, interactuar con otros personajes y completar misiones.
// Implementar habilidades y poderes para los personajes, así como sistemas de combate para enfrentarse a enemigos.
// Proporcionar una interfaz de usuario intuitiva para que los jugadores puedan interactuar con el juego y ver su progreso.
// Utilizar la creatividad y la imaginación para hacer que el juego sea divertido y emocionante para los jugadores.

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Enemigo struct {
	Id      int
	Nombre  string
	Salud   int
	Ataque  int
	Defensa int
}

type Objeto struct {
	Nombre  string
	Salud   int
	Ataque  int
	Defensa int
}

type Objetos struct {
	Objetos []Objeto
}

type Personaje struct {
	Nombre  string
	Salud   int
	Ataque  int
	Defensa int
	Objetos Objetos
	Mision  Mision
}

type Mision struct {
	Nombre     NombreMisiones
	Completado bool
}

type Misiones struct {
	Misiones []Mision
}

type NombreMisiones string

const (
	mision1 NombreMisiones = "CREAR PERSONAJE"
	mision2 NombreMisiones = "MATAR ENEMIGO"
	mision3 NombreMisiones = "RECUPERAR ESPADA"
	mision4 NombreMisiones = "RECUPERAR SALUD"
)

func crearEnemigo(enemigosMap map[string]Enemigo, id, salud, ataque, defensa int, nombre string) {
	nuevoEnemigo := Enemigo{
		Id:      id,
		Nombre:  nombre,
		Salud:   salud,
		Ataque:  ataque,
		Defensa: defensa,
	}
	enemigosMap[nuevoEnemigo.Nombre] = nuevoEnemigo
}

func asignarEstadisticas(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func crearJugador(nombre string) *Personaje {
	Personaje := &Personaje{
		Nombre:  nombre,
		Salud:   100,
		Ataque:  asignarEstadisticas(800, 1000),
		Defensa: asignarEstadisticas(1000, 2000),
		Objetos: Objetos{},
		Mision: Mision{
			Nombre:     mision1,
			Completado: true,
		},
	}
	return Personaje
}

func crearMision(misiones *[]Mision, nombre NombreMisiones) {
	nuevaMision := Mision{
		Nombre:     nombre,
		Completado: false,
	}

	*misiones = append(*misiones, nuevaMision)
}

func seleccionarEnemigoAleatorio(enemigosMap map[string]Enemigo) Enemigo {
	var enemigos []Enemigo
	for _, enemigo := range enemigosMap {
		enemigos = append(enemigos, enemigo)
	}
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(enemigos))
	return enemigos[index]
}

func enfrentarEnemigo(jugador *Personaje, enemigo Enemigo) bool {
	fmt.Printf("¡%s ha aparecido!\n", enemigo.Nombre)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("\n%s (Salud: %d, Ataque: %d, Defensa: %d) vs %s (Salud: %d, Ataque: %d, Defensa: %d)\n",
			jugador.Nombre, jugador.Salud, jugador.Ataque, jugador.Defensa,
			enemigo.Nombre, enemigo.Salud, enemigo.Ataque, enemigo.Defensa)

		fmt.Println("¿Qué deseas hacer?")
		fmt.Println("1. Atacar")
		fmt.Println("2. Retirarse")
		fmt.Print("Selecciona una opción: ")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			// Calcular el daño del jugador al enemigo
			dañoJugador := jugador.Ataque - enemigo.Defensa
			
			if dañoJugador < 0 {
				dañoJugador = 0
			}
			enemigo.Salud -= dañoJugador

			fmt.Printf("%s hace %d de daño a %s.\n", jugador.Nombre, dañoJugador, enemigo.Nombre)

			// Verificar si el enemigo ha sido derrotado
			if enemigo.Salud <= 0 {
				fmt.Printf("¡%s ha sido derrotado!\n", enemigo.Nombre)
				return true
			}

			// Calcular el daño del enemigo al jugador
			dañoEnemigo := enemigo.Ataque - jugador.Defensa
			if dañoEnemigo < 0 {
				dañoEnemigo = 0
			}
			jugador.Salud -= dañoEnemigo

			fmt.Printf("%s hace %d de daño a %s.\n", enemigo.Nombre, dañoEnemigo, jugador.Nombre)

			// Verificar si el jugador ha sido derrotado
			if jugador.Salud <= 0 {
				fmt.Printf("%s ha sido derrotado por %s. Game Over.\n", jugador.Nombre, enemigo.Nombre)
				os.Exit(0)
			}

		case "2":
			fmt.Printf("%s decide retirarse del combate.\n", jugador.Nombre)
			return false

		default:
			fmt.Println("Opción no válida. Por favor selecciona 1 o 2.")
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	EnemigosMap := make(map[string]Enemigo)
	crearEnemigo(EnemigosMap, 1, 100, asignarEstadisticas(500, 1000), asignarEstadisticas(1000, 2000), "FANTASMA")
	crearEnemigo(EnemigosMap, 2, 100, asignarEstadisticas(500, 1000), asignarEstadisticas(1000, 2000), "MAGO")
	crearEnemigo(EnemigosMap, 3, 100, asignarEstadisticas(500, 1000), asignarEstadisticas(1000, 2000), "VAMPIRO")
	crearEnemigo(EnemigosMap, 4, 100, asignarEstadisticas(500, 1000), asignarEstadisticas(1000, 2000), "BRUJA")

	var misiones []Mision
	crearMision(&misiones, mision2)
	crearMision(&misiones, mision3)
	crearMision(&misiones, mision4)

	jugador := crearJugador("Jugador")

	fmt.Printf("¡Bienvenido a la aventura, %s!\n", jugador.Nombre)

	for _, mision := range misiones {
		if !mision.Completado {
			fmt.Printf("\nComienza la misión: %s\n", mision.Nombre)
			enemigo := seleccionarEnemigoAleatorio(EnemigosMap)
			fmt.Printf("Enfrentando al enemigo para completar la misión: %s\n", enemigo.Nombre)

			if enfrentarEnemigo(jugador, enemigo) {
				fmt.Printf("¡Has derrotado a %s! Misión completada.\n", enemigo.Nombre)
				mision.Completado = true
			} else {
				fmt.Printf("Te has retirado de la misión: %s. Mejor suerte la próxima vez.\n", mision.Nombre)
			}
		}
	}

	fmt.Println("\n¡Felicidades! Has completado todas las misiones. ¡Juego terminado!")
}
