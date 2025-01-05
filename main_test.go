package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMotdHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/motd", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	motdHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected status 200; got %v", rec.Code)
	}

	expected := "Welcome to the Go MOTD Server! Have a great day!"
	if rec.Body.String() != expected {
		t.Errorf("Expected body %q; got %q", expected, rec.Body.String())
	}
}

