package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"app/url-shorter/configs"
	"app/url-shorter/pkg/response"

	"github.com/go-playground/validator/v10"
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
		validate := validator.New()
		err = validate.Struct(payload)
		if err != nil {
			response.Json(w, err.Error(), 402)
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
