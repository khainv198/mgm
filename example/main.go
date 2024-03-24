package main

import (
	"context"
	"log"

	"github.com/khainv198/mgm"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `bson:"name,omitempty"`
}

func main() {
	client, _ := mgm.New(context.Background(), "test_db", options.Client().ApplyURI("mongodb://localhost:27017"))

	book := &Book{Name: "book name"}

	client.Collection("books").Create(context.Background(), book)
	log.Print(book.ID.Hex())
}
