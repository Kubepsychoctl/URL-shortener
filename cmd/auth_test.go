package main

import (
	"app/url-shorter/internal/auth"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginSucces(t *testing.T) {
	server := httptest.NewServer(Init())
	defer server.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "j@d.ru",
		Password: "2",
	})

	response, err := http.Post(server.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, response.StatusCode)
	}
}
