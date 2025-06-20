package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
}

func main() {
	http.HandleFunc("/hello", hello)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
