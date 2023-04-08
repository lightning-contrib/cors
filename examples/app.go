package main

import (
	"github.com/go-labx/lightning"
	"github.com/lightning-contrib/cors"
)

func main() {
	app := lightning.DefaultApp()

	// Add the CORS middleware to the middleware chain
	app.Use(cors.Default())

	// Add your routes here
	app.Get("/ping", func(ctx *lightning.Context) {
		ctx.Success("hello world")
	})

	// Start the server
	app.Run()
}
