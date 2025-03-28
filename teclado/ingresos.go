package teclado

import (
	"fmt"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {
	fmt.Println("Ingrese numero: ")
	fmt.Scanln(&numero1)
	fmt.Printf("El numero ingresado es: %d \n", numero1)

}
