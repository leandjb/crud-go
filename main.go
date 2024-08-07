package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"strconv"
)

var reader *bufio.Reader

type User struct {
	id int
	username string
	email string
	age int
}

var id int
var users map[int]User


func eliminarUsuario() {
	fmt.Println("MSG: Usuario eliminado.")
	
}

func modificarUsuario() {
	fmt.Println("MSG: Usuario actualizado exitosamente.")
	
}

func listarUsuarios() {
	fmt.Println("*** Lista de Usuarios ***")
	
}

func crearUsuario() {
	fmt.Println("> Ingrese un valor para username:")
	username := readLine()
	
	fmt.Println("> Ingrese un valor para email:")
	email := readLine()
	
	fmt.Println("> Ingrese un valor para edad: ")
	age, err := strconv.Atoi(readLine())

	if err != nil {
		fmt.Println("MSG: No es posible convertir string a un entero.")
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

func main() {

	var option string
	users = make(map[int]User)
	reader = bufio.NewReader(os.Stdin)
	
	for{
		fmt.Println("-------- MENU --------")
		fmt.Println("A. Crear usuario")
		fmt.Println("B. Listar usuarios")
		fmt.Println("C. Modificar usuario")
		fmt.Println("D. Eliminar usuario")
		fmt.Println("E. Salir")
		fmt.Println("______________________")
	
		fmt.Println("> Ingresa una opcion:")
		option = readLine()	
		
		if option == "salir" || option == "e"{
			break
		}
	
		switch option {
		case "a", "crear":
			crearUsuario()
		case "b", "listar":
			listarUsuarios()
		case "c", "modificar":
			modificarUsuario()
		case "d", "eliminar":
			eliminarUsuario()
	
		default:
			fmt.Println("MSG: Opcion invalida.")
		}
	}

	fmt.Println("Adios!")
}