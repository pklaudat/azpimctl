package utils

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Request(operation, url, token string, body io.Reader) io.ReadCloser {

	req, err := http.NewRequest(operation, url, body)

	if err != nil {
		panic("Failed to create new request system.")
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(req)

	if err != nil {
		panic(fmt.Sprintf("HTTP request to system failed: %s %s", err.Error(), url))
	}

	resp.Body.Close()

	return resp.Body
}
