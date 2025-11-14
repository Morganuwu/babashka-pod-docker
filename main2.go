package main

import (
	"fmt"
	"strings"
	"time"
)

// Estructura de datos del usuariosjsjjs
// //prueba de uwu
type Usuario struct {
	ID     int
	Nombre string
	Correo string
	Edad   int
	Ciudad string
	Calle  string
	Puesto string
}

func main() {
	//  Login de acceso
	if !login() {
		fmt.Println("\nDemasiados intentos fallidos. Cerrando el sistema.")
		return
	}

	// Base de datos simulada
	usuarios := []Usuario{
		{1, "Ana Torres", "ana.torres@example.com", 29, "Ciudad de México", "Av. Reforma #123", "Administradora"},
		{2, "Luis Gómez", "luis.gomez@example.com", 35, "Guadalajara", "Calle Hidalgo #56", "Técnico de soporte"},
		{3, "María López", "maria.lopez@example.com", 31, "Monterrey", "Av. Constitución #22", "Diseñadora"},
		{4, "Carlos Rivera", "carlos.rivera@example.com", 28, "Puebla", "Calle 5 de Mayo #67", "Desarrollador"},
		{5, "Lucía Hernández", "lucia.hernandez@example.com", 40, "Querétaro", "Blvd. Bernardo Quintana #90", "Gerente de proyecto"},
	}

	fmt.Println("=======================================")
	fmt.Println("   Sistema de Búsqueda y Gestión de Usuarios")
	fmt.Println("=======================================")

	for {
		fmt.Println("\nSeleccione una opción:")
		fmt.Println("1. Buscar usuario por nombre")
		fmt.Println("2. Mostrar todos los usuarios")
		fmt.Println("3. Modificar información de un usuario")
		fmt.Println("4. Salir")
		fmt.Print("Opción: ")

		var opcion int
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			buscar(usuarios)
		case 2:
			listar(usuarios)
		case 3:
			usuarios = modificar(usuarios)
		case 4:
			fmt.Println("\nSaliendo del sistema. Hasta pronto.")
			return
		default:
			fmt.Println("Opción no válida. Intente nuevamente.")
		}
	}
}

// Función de login
func login() bool {
	const usuarioCorrecto = "morgan"
	const contrasenaCorrecta = "123"
	intentos := 0
	maxIntentos := 3

	fmt.Println("=======================================")
	fmt.Println("          Acceso al Sistema")
	fmt.Println("=======================================")

	for intentos < maxIntentos {
		var usuario, contrasena string

		fmt.Print("Usuario: ")
		fmt.Scanln(&usuario)

		fmt.Print("Contraseña: ")
		fmt.Scanln(&contrasena)

		if usuario == usuarioCorrecto && contrasena == contrasenaCorrecta {
			fmt.Println("\n✅ Acceso concedido. Bienvenido,", usuario)
			return true
		}

		intentos++
		fmt.Printf(" Credenciales incorrectas. Intento %d de %d.\n", intentos, maxIntentos)
	}

	return false
}

// Buscar usuario por nombre
func buscar(usuarios []Usuario) {
	fmt.Print("\nIngrese el nombre a buscar: ")
	var busqueda string
	fmt.Scanln(&busqueda)

	if busqueda == "" {
		fmt.Println("No se puede buscar un nombre vacío.")
		return
	}

	fmt.Println("Buscando coincidencias...")
	time.Sleep(1 * time.Second)

	encontrados := []Usuario{}
	for _, u := range usuarios {
		if strings.Contains(strings.ToLower(u.Nombre), strings.ToLower(busqueda)) {
			encontrados = append(encontrados, u)
		}
	}

	if len(encontrados) == 0 {
		fmt.Println("No se encontraron resultados.")
	} else {
		fmt.Println("\nResultados encontrados:")
		for _, u := range encontrados {
			mostrarUsuario(u)
		}
	}
}

// Mostrar todos los usuarios
func listar(usuarios []Usuario) {
	fmt.Println("\nLista completa de usuarios:")
	for _, u := range usuarios {
		mostrarUsuario(u)
	}
}

// Modificar usuario
func modificar(usuarios []Usuario) []Usuario {
	fmt.Print("\nIngrese el ID del usuario que desea modificar: ")
	var id int
	fmt.Scanln(&id)

	index := -1
	for i, u := range usuarios {
		if u.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Usuario no encontrado.")
		return usuarios
	}

	fmt.Println("\nUsuario encontrado:")
	mostrarUsuario(usuarios[index])

	fmt.Println("\nIngrese los nuevos datos (deje vacío para mantener el actual):")

	var nuevoNombre, nuevoCorreo, nuevaCiudad, nuevaCalle, nuevoPuesto string
	var nuevaEdad int

	fmt.Printf("Nombre (%s): ", usuarios[index].Nombre)
	fmt.Scanln(&nuevoNombre)
	if nuevoNombre != "" {
		usuarios[index].Nombre = nuevoNombre
	}

	fmt.Printf("Correo (%s): ", usuarios[index].Correo)
	fmt.Scanln(&nuevoCorreo)
	if nuevoCorreo != "" {
		usuarios[index].Correo = nuevoCorreo
	}

	fmt.Printf("Edad (%d): ", usuarios[index].Edad)
	fmt.Scanln(&nuevaEdad)
	if nuevaEdad != 0 {
		usuarios[index].Edad = nuevaEdad
	}

	fmt.Printf("Ciudad (%s): ", usuarios[index].Ciudad)
	fmt.Scanln(&nuevaCiudad)
	if nuevaCiudad != "" {
		usuarios[index].Ciudad = nuevaCiudad
	}

	fmt.Printf("Calle (%s): ", usuarios[index].Calle)
	fmt.Scanln(&nuevaCalle)
	if nuevaCalle != "" {
		usuarios[index].Calle = nuevaCalle
	}

	fmt.Printf("Puesto (%s): ", usuarios[index].Puesto)
	fmt.Scanln(&nuevoPuesto)
	if nuevoPuesto != "" {
		usuarios[index].Puesto = nuevoPuesto
	}

	fmt.Println("\n✅ Información actualizada correctamente.")
	return usuarios
}

// 👤 Mostrar información de un usuario
func mostrarUsuario(u Usuario) {
	fmt.Printf("\nID: %d\nNombre: %s\nCorreo: %s\nEdad: %d\nCiudad: %s\nCalle: %s\nPuesto: %s\n",
		u.ID, u.Nombre, u.Correo, u.Edad, u.Ciudad, u.Calle, u.Puesto)
}
