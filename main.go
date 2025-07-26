package main

import (
	"context"
	"log"
	"os"
	// "strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

var collection *mongo.Collection

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// Get port from environment variable
	MONGODB_URI := os.Getenv("MONGODB_URI")
	PORT := os.Getenv("PORT")

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}

	defer client.Disconnect(context.Background())

	// Check connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB: ", err)
	} else {
		log.Println("Connected to MongoDB")
	}

	// Get collection
	collection = client.Database("react_go").Collection("todos")

	// Create a new Fiber app
	app := fiber.New()

	// Define routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	// Create API route group
	api := app.Group("/api")

	// Now all routes under this group will have /api prefix
	api.Get("/todos", getTodos)
	api.Get("/todos/:id", getTodoById)
	api.Post("/todos", createTodo)
	api.Put("/todos/:id", updateTodo)
	api.Patch("/todos/:id", completeTodo)
	api.Delete("/todos/:id", deleteTodo)

	// Listen on port
	if PORT == "" {
		PORT = "8000"
	}
	log.Fatal(app.Listen("0.0.0.0:" + PORT))
}

func getTodos(c *fiber.Ctx) error {
	var todos []Todo = []Todo{}

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
			"message": "Error retrieving todos",
		})
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
				"message": "Error decoding todo",
			})
		}
		todos = append(todos, todo)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Todos retrieved successfully",
		"data":    todos,
	})
}

func getTodoById(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
			"message": "Invalid ID format",
		})
	}

	filter := bson.M{"_id": id}

	var todo Todo
	err = collection.FindOne(context.Background(), filter).Decode(&todo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).JSON(fiber.Map{
				"success": false,
				"message": "Todo not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
			"message": "Error retrieving todo",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Todo retrieved successfully",
		"data":    todo,
	})
}

func createTodo(c *fiber.Ctx) error {
	todo := new(Todo) // { id: 0, completed: false, body: "" }

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

	todo.ID = primitive.NewObjectID()

	_, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
			"message": "Error creating todo",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"data":    todo,
		"message": "Todo created",
	})
}

func updateTodo(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
			"message": "Invalid ID format",
		})
	}

	var updateData struct {
		Body string `json:"body"`
	}

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
			"message": "Cannot parse JSON",
		})
	}

	if updateData.Body == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Body is required",
		})
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"body": updateData.Body}}

	result, err := collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
			"message": "Error updating todo",
		})
	}

	if result.ModifiedCount == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Todo not found",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Todo updated",
	})
}

func completeTodo(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
			"message": "Invalid ID format",
		})
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"completed": true}}

	result, err := collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
			"message": "Error updating todo",
		})
	}

	if result.ModifiedCount == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Todo not found",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Todo completed",
	})
}

func deleteTodo(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
			"message": "Invalid ID format",
		})
	}

	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
			"message": "Error deleting todo",
		})
	}

	if result.DeletedCount == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Todo not found",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Todo deleted",
	})
}
