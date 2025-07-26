package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
    ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Completed bool               `json:"completed" bson:"completed"`
    Body      string             `json:"body" bson:"body"`
}

type TodoRequest struct {
    Body      string `json:"body" validate:"required"`
}

type TodoResponse struct {
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}
