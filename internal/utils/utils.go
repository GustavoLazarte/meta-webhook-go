package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func CreateRequest(body io.Reader) (*http.Request, error) {
	url := fmt.Sprintf("https://graph.facebook.com/v19.0/%s/messages", os.Getenv("META_NUMBER_ID"))

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Printf("Error procesando el mensaje: %v", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("META_TOKEN"))

	return req, err
}

func MapStruct(source interface{}, destination interface{}) error {
	bytes, err := json.Marshal(source)
	if err != nil {
		return fmt.Errorf("error al serializar source: %w", err)
	}

	if err := json.Unmarshal(bytes, destination); err != nil {
		return fmt.Errorf("error al deserializar en destination: %w", err)
	}

	return nil
}
