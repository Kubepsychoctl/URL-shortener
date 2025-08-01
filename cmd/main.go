package main

import (
	"fmt"
	"net/http"

	"app/url-shorter/configs"
	"app/url-shorter/internal/auth"
	"app/url-shorter/internal/link"
	"app/url-shorter/pkg/db"
	"app/url-shorter/pkg/middleware"
)

func main() {
	config := configs.LoadConfig()
	db := db.NewDb(config)
	router := http.NewServeMux()

	linkRepo := link.NewLinkRepository(db)
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: config,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepo: linkRepo,
	})
	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.Logging(router),
	}

	fmt.Println("Server is running on port 8080")
	server.ListenAndServe()
}
