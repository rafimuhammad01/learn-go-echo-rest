package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database
var client *mongo.Client
var err error

func Init() {
	godotenv.Load(".env")
	username := os.Getenv("DB_USERNAME")
	pw := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb+srv://" + username + ":" + pw+ "@cluster0.dumg5.mongodb.net/" + name+ "?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//defer client.Disconnect(ctx)


	db = client.Database("echo-rest-api")
}

func Connect() (*mongo.Database, error) {
	return db, err
}

func Disconnect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client.Disconnect(ctx)
}
