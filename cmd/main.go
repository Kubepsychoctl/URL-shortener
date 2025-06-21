package main

import (
	"fmt"
	"net/http"

	"app/url-shorter/internal/hello"
)

func main() {
	router := http.NewServeMux()
	hello.NewHelloHandler(router)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server is running on port 8080")
	server.ListenAndServe()
}
