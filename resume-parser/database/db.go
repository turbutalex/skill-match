package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ResumeCollection *mongo.Collection

func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017") // Change as needed
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the database and collection
	ResumeCollection = client.Database("resume_parser_db").Collection("resumes")
	log.Println("Connected to MongoDB!")
}
