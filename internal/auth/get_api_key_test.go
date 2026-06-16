package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	h := http.Header{}
	_, err := GetAPIKey(h)
	if err == ErrNoAuthHeaderIncluded {
		t.Errorf("should return error no auth header included")
	}
}
