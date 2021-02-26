package main

import (
	"log"

	"github.com/xsama/twitsam/bd"
	"github.com/xsama/twitsam/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexion en la bd")
		return
	}
	handlers.Manejadores()
}
