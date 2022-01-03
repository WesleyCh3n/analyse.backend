package main

import (
	"flag"
	"server/handlers"

	"github.com/gofiber/fiber/v2"
)

var (
	port       = flag.String("port", ":3000", "Port to listen on")
	prod       = flag.Bool("prod", false, "Enable prefork in Production")
	staticFile = flag.String("d", "./static/build/", "static web folder")
)

func setupRoutes(app *fiber.App) {
	apiGroup := app.Group("/api")
	apiGroup.Get("/ping", handlers.Pong)
	apiGroup.Post("/upload", handlers.UploadFile)

	fileGroup := app.Group("/file")
	fileGroup.Static("/csv", "./file/csv/")
}

func NewServer() *fiber.App {
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024, // 100 mb
	})

	setupRoutes(app)

	app.Static("/", *staticFile)

	return app
}

func main() {
	flag.Parse()

	app := NewServer()
	app.Listen(*port)
}
