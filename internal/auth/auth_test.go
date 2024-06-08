package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_NoAuthHeader(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatal("expected an error but got none")
	}

	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected %v but got %v", ErrNoAuthHeaderIncluded, err)
	}
}

func TestGetAPIKey_ValidAuthHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization1", "ApiKey testkey123")

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error but got %v", err)
	}

	expectedAPIKey := "testkey123"
	if apiKey != expectedAPIKey {
		t.Fatalf("expected %v but got %v", expectedAPIKey, apiKey)
	}
}
