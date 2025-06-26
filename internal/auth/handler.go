package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"

	"app/url-shorter/configs"
	"app/url-shorter/pkg/response"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload LoginRequest
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			response.Json(w, err.Error(), 402)
			return
		}
		if payload.Email == "" {
			response.Json(w, "Email is required", 402)
			return
		}
		if !validateEmail(payload.Email) {
			response.Json(w, "Invalid email", 402)
			return
		}
		if payload.Password == "" {
			response.Json(w, "Password is required", 402)
			return
		}
		fmt.Println(payload)
		res := LoginResponse{
			Token: "1234567890",
		}
		response.Json(w, res, 200)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
