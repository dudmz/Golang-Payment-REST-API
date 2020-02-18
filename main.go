package main

import (
	"github.com/carlosdamazio/Stone-REST-API/app"
)

func main() {
	application := &app.App{}
	application.Initialize()
	application.Run(":8080")
}
