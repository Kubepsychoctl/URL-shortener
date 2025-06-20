package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	NewHelloHandler(router)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server is running on port 8080")
	server.ListenAndServe()
}
