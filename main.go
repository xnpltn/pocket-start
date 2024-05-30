package main

import (
	"log"

	"github.com/xnpltn/pocket-start/application"
)

func main() {
	app := application.NewApp()

	log.Fatal(app.Start())
}
