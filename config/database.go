package config

import (
    "context"
    "log"
    "os"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
    MONGODB_URI := os.Getenv("MONGODB_URI")
    if MONGODB_URI == "" {
        log.Fatal("MONGODB_URI environment variable is not set")
    }

    clientOptions := options.Client().ApplyURI(MONGODB_URI)
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal("Error connecting to MongoDB: ", err)
    }

    // Check connection
    err = client.Ping(context.Background(), nil)
    if err != nil {
        log.Fatal("Error pinging MongoDB: ", err)
    }

    log.Println("Connected to MongoDB")
    DB = client.Database("react_go")
}

func GetCollection(collectionName string) *mongo.Collection {
    return DB.Collection(collectionName)
}
