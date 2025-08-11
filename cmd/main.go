package main

import (
	"fmt"
	"net/http"

	"app/url-shorter/configs"
	"app/url-shorter/internal/auth"
	"app/url-shorter/internal/link"
	"app/url-shorter/internal/stat"
	"app/url-shorter/internal/user"
	"app/url-shorter/pkg/db"
	"app/url-shorter/pkg/event"
	"app/url-shorter/pkg/middleware"
)

func Init() http.Handler {
	config := configs.LoadConfig()
	db := db.NewDb(config)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	// Repositories
	linkRepo := link.NewLinkRepository(db)
	userRepo := user.NewUserRepository(db)
	statRepo := stat.NewStatRepository(db)

	// Services
	authService := auth.NewUserService(userRepo)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus: eventBus,
		StatRepo: statRepo,
	})

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      config,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepo: linkRepo,
		EventBus: eventBus,
		Config:   config,
	})
	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepo: statRepo,
		Config:   config,
	})

	go statService.AddClick()

	// Middleware
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)
	return stack(router)
}

func main() {
	app := Init()
	server := http.Server{
		Addr:    ":8080",
		Handler: app,
	}
	fmt.Println("Server is running on port 8080")
	server.ListenAndServe()
}
