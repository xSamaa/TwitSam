package main

import (
	"log"

	"github.com/xSamaa/twitsam/bd"
	"github.com/xSamaa/twitsam/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexion en la bd")
		return
	}
	handlers.Manejadores()
}
