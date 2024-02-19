package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApikey(header http.Header) (string, error) {

	api_key := header.Get("Authorization")
	if api_key == "" {
		return "", errors.New("apikey is required.")
	}

	vals := strings.Split(api_key, " ")

	if len(vals) != 2 || vals[0] != "Apikey" {
		return "", errors.New("invalid apikey format.")
	}

	return vals[1], nil

}
