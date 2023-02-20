package main

import "fmt"

// App, the struct which contains things like pointers to db connections
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up App")
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
