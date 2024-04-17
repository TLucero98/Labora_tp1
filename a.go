package main

import (
	"fmt"
	"os"
)

var (
	AdminUsername = "admin"
	AdminPassword = "admin"
	User1         = "User1"
	User1Password = "User1"
)
var (
	AdminActions = []string{"Crear archivo", "Eliminar archivo"}
)

var (
	Files = []string{}
)

func main() {
	for {
		fmt.Println("Bienvenido")
		fmt.Println("1. Iniciar como admin")
		fmt.Println("2. Iniciar como user")
		fmt.Println("3. Salir")

		var choice int
		fmt.Println("Ingrese su opcion: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			adminlogin()
		case 2:
			userlogin()
		case 3:
			fmt.Println("Adios")
			os.Exit(0)
		default:
			fmt.Println("Seleccine una opcion valida")

		}

	}
}

func adminlogin() {
	fmt.Println("Ingrese las credenciales de admin")
	fmt.Println("Usuario: ")
	var usuario string
	fmt.Scanln(&usuario)
	fmt.Println("Contraseña: ")
	var password string
	fmt.Scanln(&password)

	if usuario == AdminUsername && password == AdminPassword {
		fmt.Println("Inicio de sesion de admin completo")
		adminMenu()
	} else {
		fmt.Println("Credenciales incorrectas, intente de nuevo")

	}
}

func userlogin() {
	fmt.Println("Ingrese las credenciales de usuario ")
	fmt.Println("Usuario: ")
	var usuario string
	fmt.Scanln(&usuario)
	var password string
	fmt.Scanln(&password)

	if usuario == User1 && password == User1Password {
		fmt.Println("Inicio de sesion de usuario completo")
		userMenu()
	} else {
		fmt.Println("Credenciales incorrectas, intente de nuevo")

	}

}

func adminMenu() {
	for {
		fmt.Println("Menu de admin")
		for i, action := range AdminActions {
			fmt.Printf("%d. %s\n", i+1, action)
		}
		fmt.Println("0. Cerrar sesión")

		var choice int
		fmt.Println("Seleccione su opcion: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Crear archivo")
			createFile()
		case 2:
			fmt.Println("Eliminar archivo")
			deleteFile()
		case 0:
			fmt.Println("Cerrar sesión de admin")
			return
		default:
			fmt.Println("Seleccione una opcion valida")

		}

	}

}

func createFile() {
	Files = append(Files, fmt.Sprintf("file %d", len(Files)+1))
	fmt.Println("Archivo creado exitosamente")
}

func deleteFile() {
	if len(Files) == 0 {
		fmt.Println("No hay archivos para eliminar")
		return
	}
	Files = Files[:len(Files)-1]
	fmt.Println("Archivo eliminado exitosamente")
}

func userMenu() {

}
