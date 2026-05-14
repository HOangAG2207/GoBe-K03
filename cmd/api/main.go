package main

import "github.com/HOangAG2207/GoBe-K03/internal/config"

func main() {
	app := config.NewEngine()
	if err := app.Start(); err != nil {
		panic(err)
	}
}
