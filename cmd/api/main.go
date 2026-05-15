package main

import "github.com/HOangAG2207/GoBe-K03/internal/config"

// @title GoBe-K03 API
// @version 1.0
// @description This is an API for the GoBe-K03 application.
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /
func main() {
	cfg := config.Load()
	app := config.NewEngine(cfg)
	if err := app.Start(); err != nil {
		panic(err)
	}
}
