package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeyGoodHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey super-duper-secret-key")
	str, err := GetAPIKey(headers)
	wantStr := "super-duper-secret-key"
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if str != wantStr {
		t.Fatalf("expected: %v, got %v", wantStr, str)
	}
}

func TestGetAPIKeyWrongPrefix(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApeKey super-sneaky-shadowy-key")
	_, err := GetAPIKey(headers)
	if err == nil {
		t.Errorf("expected error but didn't get one")
	}
}

func TestGetAPIKeyMissingKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "unknown-key")
	_, err := GetAPIKey(headers)
	if err == nil {
		t.Errorf("expected error but didn't get one")
	}
}

func TestGetAPIKeyEmptyHeader(t *testing.T) {
	headers := http.Header{}
	_, err := GetAPIKey(headers)
	if err == nil {
		t.Errorf("expected error but didn't get one")
	}
}
