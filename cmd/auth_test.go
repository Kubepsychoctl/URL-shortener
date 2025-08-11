package main

import (
	"app/url-shorter/internal/auth"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginSucces(t *testing.T) {
	server := httptest.NewServer(Init())
	defer server.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "j@d.ru",
		Password: "1",
	})

	response, err := http.Post(server.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, response.StatusCode)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	var resData auth.LoginResponse
	err = json.Unmarshal(body, &resData)
	if err != nil {
		t.Fatal(err)
	}
	if resData.Token == "" {
		t.Fatal("Token is nil")
	}

}

func TestLoginFail(t *testing.T) {
	server := httptest.NewServer(Init())
	defer server.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "j@d.ru",
		Password: "123456",
	})

	response, err := http.Post(server.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected %d, got %d", http.StatusUnauthorized, response.StatusCode)
	}
}
