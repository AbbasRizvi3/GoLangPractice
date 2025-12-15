package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var client *mongo.Client

type User struct {
	Name     string `json:"userName"`
	Age      int    `json:"userAge"`
	Email    string `json:"userEmail"`
	Password string `json:"userPassword"`
}

// create multiple router with concurreny for multiple requests

type Routes struct {
	handlers map[string]http.HandlerFunc
}

func NewRoutes() *Routes {
	routes := Routes{
		handlers: make(map[string]http.HandlerFunc),
	}
	return &routes
}

func (r *Routes) AddRoute(route string, handler http.HandlerFunc) {
	r.handlers[route] = handler
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User's page")
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request:", r.Method, r.URL.Path)
		next(w, r) // call the next handler
	}
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "secret-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func (r *Routes) ServeHTTP(h http.ResponseWriter, w *http.Request) {
	for pattern, handler := range r.handlers {
		if matched, _ := path.Match(pattern, w.URL.Path); matched {
			handler.ServeHTTP(h, w)
			return
		}
	}
	http.NotFound(h, w)
}

func handleDashboard(w http.ResponseWriter, h *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("Error loading env")
	}
	uri := os.Getenv("MONGODB_URI")
	docs := "www.mongodb.com/docs/drivers/go/current/"
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " + docs +
			"usage-examples/#environment-variable")
	}
	client, err = mongo.Connect(options.Client().
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

	routes := NewRoutes()
	routes.AddRoute("/", loggingMiddleware(authMiddleware(handleDashboard)))
	routes.AddRoute("/users", handleUsers)

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", routes))

}
