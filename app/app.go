package app

import (
	"fmt"
)

type App struct {
	Router string
	DB	   string
}

func (a *App) Initialize() {
	fmt.Println(a.DB, a.Router)
	fmt.Println("This will be the initialization of the REST API...")
}