package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const client string

type User struct {
	Name     string
	Age      int
	Email    string
	Password string
}

func HandleSignUp(h http.ResponseWriter, r *http.Request) {

	h.WriteHeader(200)

}

func main() {

	envs := godotenv.Load()
	if envs != nil {
		fmt.Errorf("Error loading env")
	}
	uri := os.Getenv("MONGODB_URI")
	docs := "www.mongodb.com/docs/drivers/go/current/"
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " + docs +
			"usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	} else {
		fmt.Print("Connected")
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// http.HandleFunc("/signup")

}
