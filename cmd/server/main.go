package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/QiaoLin22/Golang-production-1/internal/transport/http"
)

// App, the struct which contains things like pointers to db connections
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up App")
	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up REST API")
		fmt.Println(err)
	}
}
