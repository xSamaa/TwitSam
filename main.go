package main

import (
	"log"

	"github.com/xSamaa/TwitSam/bd"
	"github.com/xSamaa/TwitSam/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la base de datos")
		return
	}
	handlers.Manejadores()
}
