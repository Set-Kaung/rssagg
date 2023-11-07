package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extract api_key from http request header
func GetAPIKey(header http.Header) (string, error) {
	val := header.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}
	vals := strings.Split(val, ":")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "APIKey" {
		return "", errors.New("malformed first part of error")
	}
	return vals[1], nil
}
