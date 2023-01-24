package main

import (
	"log"

	"github.com/jestebanrods/edu-golang/dispatch/internal/infrastructure/di"
)

func main() {
	app, err := di.InitializeApp()
	if err != nil {
		log.Fatal("fallo la construccion del server")
	}

	err = app.Run()
	if err != nil {
		log.Fatal("fallo la ejecucion del server")
	}
}
