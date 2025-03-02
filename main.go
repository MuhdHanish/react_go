package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	// Create a new Fiber app
	app := fiber.New()

	// Tdod's
	todos := []Todo{}

	// Define routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{} // { id: 0, completed: false, body: "" }

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Body is required",
			})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})

	// Listen on port 3000
	log.Fatal(app.Listen(":3000"))
}
