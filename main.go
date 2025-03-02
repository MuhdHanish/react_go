package main

import (
	"log"
	
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber app
	app := fiber.New()

	// Listen on port 3000
	log.Fatal(app.Listen(":3000"))
}