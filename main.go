package main

import (
	"fmt"
	"strings"

	"github.com/punigonzalez/go-pruebas/models"
)

func sumar(a int, b int) int {
	return a + b
}

func main() {

	// imprimimo con println
	fmt.Println("¡Hola, Mundo!")

	// guardamos resultado de una funcion sumar en variable resultado
	resultado := sumar(2, 4)
	fmt.Println("Resultado:", resultado)

	// declaramos variable sin tipo con :=
	edad := 13

	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad :/")
	}

	// declaramos un slice con 3 valores
	frutas := []string{"Manzana", "Banana", "Naranja"}
	fmt.Println("Array de frutas:", frutas)
	// sumamos un valor al slice frutas
	frutas = append(frutas, "Pera")
	fmt.Println("Slice de frutas:", frutas)

	// creamos una nueva persona del struct Persona
	persona1 := models.Persona{
		Nombre: "Juan",
		Edad:   30,
	}
	fmt.Println("Nombre:", persona1.Nombre, "Edad:", persona1.Edad)

	var conferenceName = "go conference"
	fmt.Println(conferenceName)

	var tickcets = 50
	fmt.Println(tickcets, "entradas")
	tickcets = 30

	// usamos Printf para imprimir texto y pasamos %v para reemplazar por las variables, que deben estar en orden despues del ""
	fmt.Printf("Ahora tenemos una cantidad de %v entradas\n", tickcets)

	fmt.Printf("Soy %v y tengo %v entradas, quien se para de mano?\n", persona1.Nombre, tickcets)

	// creamos una variable username para saber el valor del usuario
	var username string
	// usamos el metodo Scan y le pasamos como parametro &variable, & es el puntero a direccion de variable con su valor
	fmt.Println("Ingresar usuario")
	fmt.Scan(&username)

	fmt.Printf("El usuario %v quiere tener las %v entradas que posee %v \n", username, tickcets, persona1.Nombre)

	var booking = [...]string{"Pedro", "Mariano"}
	fmt.Println(booking)

	var dineroCuenta = 5000

	ptr := &dineroCuenta

	fmt.Println(dineroCuenta)
	fmt.Println("Ptr:", *ptr)

	saludo := "Hola, "
	concatenado := saludo + username

	enMayuscula := strings.ToLower(concatenado)
	fmt.Println(enMayuscula)

}
