package main

import (
	"log"
	"strconv"

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

	// Define routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	// Tdod's
	todos := []Todo{}

	// Create API route group
	api := app.Group("/api")

	// Now all routes under this group will have /api prefix
	// Get all todos
	api.Get("/todos", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "Todos retrieved successfully",
			"data":    todos,
		})
	})

	// Get a single todo
	api.Get("/todos/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
				"message": "Invalid ID format",
			})
		}

		for _, todo := range todos {
			if todo.ID == id {
				return c.JSON(fiber.Map{
					"success": true,
					"data":    todo,
					"message": "Todo retrieved successfully",
				})
			}
		}

		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Todo not found",
		})
	})

	// Create a new todo
	api.Post("/todos", func(c *fiber.Ctx) error {
		todo := &Todo{} // { id: 0, completed: false, body: "" }

		if err := c.BodyParser(todo); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
				"message": "Cannot parse JSON",
			})
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": "Body is required",
			})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(fiber.Map{
			"success": true,
			"data":    todo,
			"message": "Todo created",
		})
	})

	// Update a todo
	api.Put("/todos/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
				"message": "Invalid ID format",
			})
		}

		for i, todo := range todos {
			if todo.ID == id {
				if err := c.BodyParser(&todos[i]); err != nil {
					return c.Status(400).JSON(fiber.Map{
						"success": false,
						"error":   err.Error(),
						"message": "Cannot parse JSON",
					})
				}
				todos[i].ID = id
				return c.JSON(fiber.Map{
					"success": true,
					"data":    todos[i],
					"message": "Todo updated",
				})
			}
		}

		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Todo not found",
		})
	})

	// Complete a todo
	api.Patch("/todos/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
				"message": "Invalid ID format",
			})
		}

		for i, todo := range todos {
			if todo.ID == id {
				todos[i].Completed = !todos[i].Completed
				return c.JSON(fiber.Map{
					"success": true,
					"data":    todos[i],
					"message": "Todo completed",
				})
			}
		}

		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Todo not found",
		})
	})

	// Delete a todo
	api.Delete("/todos/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
				"message": "Invalid ID format",
			})
		}

		for i, todo := range todos {
			if todo.ID == id {
				// Remove the todo from the slice
				todos = append(todos[:i], todos[i+1:]...)

				return c.Status(200).JSON(fiber.Map{
					"success": true,
					"message": "Todo deleted",
				})
			}
		}

		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Todo not found",
		})
	})

	// Listen on port 8000
	log.Fatal(app.Listen(":8000"))
}
