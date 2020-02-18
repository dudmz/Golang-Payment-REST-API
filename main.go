package main

import (
	"github.com/carlosdamazio/Stone-REST-API/app"
)

func main() {
	application := &app.App{Router:"Router", DB:"Database Client"}
	application.Initialize()
}
