package main

import "github.com/HOangAG2207/GoBe-K03/internal/config"

func main() {
	cfg := config.Load()
	app := config.NewEngine(cfg)
	if err := app.Start(); err != nil {
		panic(err)
	}
}
