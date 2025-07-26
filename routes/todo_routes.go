package routes

import (
    "github.com/MuhdHanish/react_go/handlers"

    "github.com/gofiber/fiber/v2"
)

func SetupTodoRoutes(api fiber.Router) {
    api.Get("/todos", handlers.GetTodos)
    api.Get("/todos/:id", handlers.GetTodoById)
    api.Post("/todos", handlers.CreateTodo)
    api.Put("/todos/:id", handlers.UpdateTodo)
    api.Patch("/todos/:id", handlers.CompleteTodo)
    api.Delete("/todos/:id", handlers.DeleteTodo)
}
