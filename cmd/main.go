package main

import (
	"fmt"
	"net/http"

	// "app/url-shorter/configs"
	"app/url-shorter/internal/auth"
)

func main() {
	// config := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server is running on port 8080")
	server.ListenAndServe()
}
