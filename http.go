package ripego

import (
	"io"
	"net/http"
)

func httpsGet(target string) (*[]byte, error) {
	resp, err := http.Get("https://" + target)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &responseBody, nil
}
