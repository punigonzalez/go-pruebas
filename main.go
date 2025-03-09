package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("¡Hola, Mundo!")
	router := gin.Default()

	router.GET("/saludo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"mensaje": "Hola, bienvenido a tu app con Gin!"})
	})

	router.Run(":8080")
}
