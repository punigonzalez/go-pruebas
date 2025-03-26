package main

import "fmt"

func main() {

	var nombre string = "santiago"
	puntero := &nombre
	nombre = "Santiaguito"
	fmt.Println(*puntero)

}
