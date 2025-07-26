package handlers

import (
    "context"
    "github.com/MuhdHanish/react_go/models"
    "github.com/MuhdHanish/react_go/utils"

    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection 

func SetCollection(col *mongo.Collection) {
    collection = col
}

func GetTodos(c *fiber.Ctx) error {
    var todos []models.Todo = []models.Todo{}

    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        return utils.ErrorResponse(c, 500, "Error retrieving todos", err)
    }
    defer cursor.Close(context.Background())

    for cursor.Next(context.Background()) {
        var todo models.Todo
        if err := cursor.Decode(&todo); err != nil {
            return utils.ErrorResponse(c, 500, "Error decoding todo", err)
        }
        todos = append(todos, todo)
    }

    return utils.SuccessResponse(c, 200, "Todos retrieved successfully", todos)
}

func GetTodoById(c *fiber.Ctx) error {
    id, err := primitive.ObjectIDFromHex(c.Params("id"))
    if err != nil {
        return utils.ErrorResponse(c, 400, "Invalid ID format", err)
    }

    var todo models.Todo
    err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&todo)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return utils.ErrorResponse(c, 404, "Todo not found", nil)
        }
        return utils.ErrorResponse(c, 500, "Error retrieving todo", err)
    }

    return utils.SuccessResponse(c, 200, "Todo retrieved successfully", todo)
}

func CreateTodo(c *fiber.Ctx) error {
    var req models.TodoRequest
    if err := c.BodyParser(&req); err != nil {
        return utils.ErrorResponse(c, 400, "Cannot parse JSON", err)
    }

    if req.Body == "" {
        return utils.ErrorResponse(c, 400, "Body is required", nil)
    }

    todo := models.Todo{
        ID:        primitive.NewObjectID(),
        Body:      req.Body,
        Completed: false,
    }

    _, err := collection.InsertOne(context.Background(), todo)
    if err != nil {
        return utils.ErrorResponse(c, 500, "Error creating todo", err)
    }

    return utils.SuccessResponse(c, 201, "Todo created", todo)
}

func UpdateTodo(c *fiber.Ctx) error {
    id, err := primitive.ObjectIDFromHex(c.Params("id"))
    if err != nil {
        return utils.ErrorResponse(c, 400, "Invalid ID format", err)
    }

    var req models.TodoRequest
    if err := c.BodyParser(&req); err != nil {
        return utils.ErrorResponse(c, 400, "Cannot parse JSON", err)
    }

    if req.Body == "" {
        return utils.ErrorResponse(c, 400, "Body is required", nil)
    }

    filter := bson.M{"_id": id}
    update := bson.M{"$set": bson.M{
        "body":      req.Body,
        "completed": false,
    }}

    result, err := collection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        return utils.ErrorResponse(c, 500, "Error updating todo", err)
    }

    if result.ModifiedCount == 0 {
        return utils.ErrorResponse(c, 404, "Todo not found", nil)
    }

    return utils.SuccessResponse(c, 200, "Todo updated", nil)
}

func CompleteTodo(c *fiber.Ctx) error {
    id, err := primitive.ObjectIDFromHex(c.Params("id"))
    if err != nil {
        return utils.ErrorResponse(c, 400, "Invalid ID format", err)
    }

    filter := bson.M{"_id": id}
    update := bson.M{"$set": bson.M{"completed": true}}

    result, err := collection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        return utils.ErrorResponse(c, 500, "Error updating todo", err)
    }

    if result.ModifiedCount == 0 {
        return utils.ErrorResponse(c, 404, "Todo not found", nil)
    }

    return utils.SuccessResponse(c, 200, "Todo completed", nil)
}

func DeleteTodo(c *fiber.Ctx) error {
    id, err := primitive.ObjectIDFromHex(c.Params("id"))
    if err != nil {
        return utils.ErrorResponse(c, 400, "Invalid ID format", err)
    }

    result, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
    if err != nil {
        return utils.ErrorResponse(c, 500, "Error deleting todo", err)
    }

    if result.DeletedCount == 0 {
        return utils.ErrorResponse(c, 404, "Todo not found", nil)
    }

    return utils.SuccessResponse(c, 200, "Todo deleted", nil)
}
