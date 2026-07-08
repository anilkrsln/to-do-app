package main

import (
	"backend/database"
	"backend/handlers"
	"backend/middleware"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	database.Connect()

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/auth/register", handlers.Register)
	mux.HandleFunc("POST /api/v1/auth/login", handlers.Login)

	mux.Handle("GET /api/v1/todos", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetTodos)))
	mux.Handle("POST /api/v1/todos", middleware.AuthMiddleware(http.HandlerFunc(handlers.CreateTodo)))
	mux.Handle("PUT /api/v1/todos/{id}", middleware.AuthMiddleware(http.HandlerFunc(handlers.UpdateTodo)))
	mux.Handle("DELETE /api/v1/todos/{id}", middleware.AuthMiddleware(http.HandlerFunc(handlers.DeleteTodo)))

	handler := middleware.CORSMiddleware(mux)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
