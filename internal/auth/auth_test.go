package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_ValidAuthHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-api-key")
	expectedAPIKey := "my-api-key"
	actualAPIKey, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if actualAPIKey != expectedAPIKey {
		t.Errorf("Expected API key to be %q, got %q", expectedAPIKey, actualAPIKey)
	}
}

func TestGetAPIKey_ValidAuthHeaderWithSpace(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-api-key ")
	expectedAPIKey := "my-api-key"
	actualAPIKey, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if actualAPIKey != expectedAPIKey {
		t.Errorf("Expected API key to be %q, got %q", expectedAPIKey, actualAPIKey)
	}
}
