package main

import (
	"github.com/MuhdHanish/react_go/config"
	"github.com/MuhdHanish/react_go/handlers"
	"github.com/MuhdHanish/react_go/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// Connect to database
	config.ConnectDB()

	// Initialize collection AFTER database connection
	collection := config.GetCollection("todos")
	handlers.SetCollection(collection)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(500).JSON(fiber.Map{
				"success": false,
				"message": "Internal Server Error",
				"error":   err.Error(),
			})
		},
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Todo API is running!",
			"version": "1.0.0",
		})
	})

	// API routes
	api := app.Group("/api")
	routes.SetupTodoRoutes(api)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen("0.0.0.0:" + port))
}
