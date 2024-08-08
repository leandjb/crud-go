package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	id       int
	username string
	email    string
	age      int
}

var id int
var users map[int]User

func eliminarUsuario() {
	clearConsole()

	fmt.Println("> Ingrese ID del usuario a eliminar:")
	id, err := strconv.Atoi(readLine())

	if err != nil {
		panic("MSG: No es posible convertir string a un entero.")
	}

	if _, ok := users[id]; ok {
		delete(users, id)
	} else {
			fmt.Println("> ID no encontrado.")
	}

	fmt.Println("MSG: Usuario eliminado.")
}

func modificarUsuario() {
	clearConsole()
	fmt.Println("> Ingrese ID del usuario a modificar:")
	id, err := strconv.Atoi(readLine())
	if err != nil {
		panic("MSG: No es posible convertir string a un entero.")
	}

	if _, ok := users[id]; ok {
		fmt.Println("> Ingrese nuevo username:")
		usernameNew := readLine()
		usernameNew = strings.ToUpper(usernameNew)

		fmt.Println("> Ingrese nuevo email:")
		emailNew := readLine()
		emailNew = strings.ToLower(emailNew)

		fmt.Println("> Ingrese una nueva edad: ")
		ageNew, err := strconv.Atoi(readLine())
		if err != nil {
			panic("MSG: No es posible convertir string a un entero.")
		}

		user := User{id, usernameNew, emailNew, ageNew}
		users[id] = user
	} else {
		fmt.Println("> ID no encontrado.")
	}



	fmt.Println("MSG: Usuario actualizado exitosamente.")
}

func listarUsuarios() {
	clearConsole()
	fmt.Println("*** Lista de Usuarios ***")

	for id, user := range users {
		fmt.Println(id, "-", user.username)
	}

	fmt.Println("\n")

}

func crearUsuario() {
	clearConsole()
	fmt.Println("> Ingrese un valor para username:")
	username := readLine()
	username = strings.ToUpper(username)

	fmt.Println("> Ingrese un valor para email:")
	email := readLine()
	email = strings.ToLower(email)

	fmt.Println("> Ingrese un valor para edad: ")
	age, err := strconv.Atoi(readLine())

	if err != nil {
		panic("MSG: No es posible convertir string a un entero.")
	}

	id++
	user := User{id, username, email, age}
	users[id] = user

	fmt.Println("MSG: Usuario creado exitosamente.")
	fmt.Println(users[id])
}

func readLine() string {

	if option, err := reader.ReadString('\n'); err != nil {
		panic("MSG: No es posible obtener el valor.")
	} else {
		option = strings.ToLower(option)
		return strings.TrimSpace(option)
	}
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {

	var option string
	users = make(map[int]User)
	reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("-------- MENU --------")
		fmt.Println("1. Crear usuario")
		fmt.Println("2. Listar usuarios")
		fmt.Println("3. Modificar usuario")
		fmt.Println("4. Eliminar usuario")
		fmt.Println("5. Salir")
		fmt.Println("______________________")

		fmt.Println("> Ingresa una opcion:")
		option = readLine()

		if option == "salir" || option == "5" {
			break
		}

		switch option {
		case "1", "crear":
			crearUsuario()
		case "2", "listar":
			listarUsuarios()
		case "3", "modificar":
			modificarUsuario()
		case "4", "eliminar":
			eliminarUsuario()

		default:
			fmt.Println("MSG: Opcion invalida.")
		}
	}

	fmt.Println("Adios!")
}
