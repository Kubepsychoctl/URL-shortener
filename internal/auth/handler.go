package auth

import (
	"fmt"
	"net/http"

	"app/url-shorter/configs"
	"app/url-shorter/pkg/request"
	"app/url-shorter/pkg/response"
)

type AuthHandlerDeps struct {
	*configs.Config
	AuthService *AuthService
}

type AuthHandler struct {
	*configs.Config
	AuthService *AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		email, err := handler.AuthService.Login(body.Email, body.Password)
		fmt.Println(email, err)
		data := LoginResponse{
			Token: "1234567890",
		}
		response.Json(w, data, 200)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		handler.AuthService.Register(body.Email, body.Password, body.Name)
	}
}
