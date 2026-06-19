package main

import (
	"qr-api/handlers"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

// Inicializa Fiber, registra las rutas y levanta el servidor.
func main() {

	app := fiber.New()

	app.Use(cors.New())

	app.Post("/qr", handlers.QRHandler)

	// Levantamos el servidor en el puerto 8080
	app.Listen(":8080")
}
