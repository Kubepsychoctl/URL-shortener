package main

import (
	"fmt"
	"net/http"

	"app/url-shorter/configs"
	"app/url-shorter/internal/auth"
	"app/url-shorter/internal/link"
	"app/url-shorter/pkg/db"
)

func main() {
	config := configs.LoadConfig()
	_ = db.NewDb(config)
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: config,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{})
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server is running on port 8080")
	server.ListenAndServe()
}
