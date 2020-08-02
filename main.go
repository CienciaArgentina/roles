package main

import "github.com/CienciaArgentina/roles/internal/app"

func main() {
	app := app.Build()
	app.Start()
}
