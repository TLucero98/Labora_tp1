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
	userActions = []string{"Ver archivos"}
)

var (
	Files    = []string{}
	Comments = []string{}
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
	//filename := fmt.Sprintf("texto_%s.txt", time.Now().Format("2006-01-02_15-04-05"))
	fmt.Println("ingrese el nombre del archivo a crear: ")
	var filename string
	fmt.Scanln(&filename)
	Files = append(Files, filename)

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error al crear archivo: %v\n", err)
		return

	}
	defer file.Close()
	for _, Files := range Files {
		_, err := file.WriteString(fmt.Sprintf("%s\n", Files))
		if err != nil {
			fmt.Println("Error al escribir el archivo: %v\n", err)
			return
		}
	}
	fmt.Printf("Archivo %s creado exitosamente.\n", filename)

	//contador de caracteres
	letterCount := len(filename)
	fmt.Println("cantidad de lentras en el nombre edl archivo: ", letterCount)
}

func deleteFile() {
	if len(Files) == 0 {
		fmt.Println("No hay archivos para eliminar")
		return
	}
	Files = Files[:len(Files)-1]
	fmt.Println("Archivo eliminado exitosamente")
	for i := 0; i < len(Files); i++ {
		fmt.Println(Files[i])

	}
}

func userMenu() {
	fmt.Println("Menu de usuario")
	for i, action := range userActions {
		fmt.Printf("%d. %s\n", i+1, action)
	}
	fmt.Println("0. Cerrar sesión")

	var choice int
	fmt.Println("Seleccione su opcion: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Println("Ver archivos")
		fmt.Println(Files)
		return
	case 0:
		fmt.Println("Cerrar sesión de usuario")
		return
	default:
		fmt.Println("Seleccione una opcion valida")

	}

}
