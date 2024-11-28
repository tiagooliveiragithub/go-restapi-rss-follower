package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey returns the API key from the headers
// Example:
// Authorization: ApiKey {insert-api-key-here}
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("missing Authorization header")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid Authorization header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of Authorization header")
	}

	return vals[1], nil
}
