package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

// SendToNode envía el payload a la API de Node.
func SendToNode(payload interface{}) ([]byte, error) {

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// Buscamos la variable de entorno que inyecta Docker.
	nodeURL := os.Getenv("NODE_URL")
	if nodeURL == "" {
		nodeURL = "http://localhost:3000/statistics"
	}

	resp, err := http.Post(
		nodeURL,
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
