package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	os2 "os"
)

const openaiAPIURL = "https://api.openai.com/v1/chat/completions"

// Solicitud para OpenAI (Modelo Chat)
type OpenAIRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"` // Puede ser "user" o "assistant"
		Content string `json:"content"`
	} `json:"messages"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
}

// Respuesta de OpenAI
type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// Estructura para manejar errores de OpenAI
type OpenAIError struct {
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error"`
}

type OpenAIService struct{}

// Manejador para responder a las solicitudes del cliente
func (os *OpenAIService) CreateRequest(question string) (*string, error) {
	// Crear mensaje para OpenAI
	messages := []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}{
		{"user", question},
	}

	// Crear solicitud para OpenAI
	requestBody, err := json.Marshal(OpenAIRequest{
		Model:       "gpt-4o-mini-2024-07-18",
		Messages:    messages,
		MaxTokens:   100,
		Temperature: 0.7,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", openaiAPIURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os2.Getenv("OPEN_API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Imprimir la respuesta completa para depuración
	fmt.Println("Respuesta de OpenAI:", string(body))

	var openAIResp OpenAIResponse
	var openAIError OpenAIError

	// Intentar deserializar como respuesta normal
	if err := json.Unmarshal(body, &openAIResp); err == nil && len(openAIResp.Choices) > 0 {
		response := openAIResp.Choices[0].Message.Content
		return &response, nil
	}

	// Si no es exitoso, intenta deserializar como un error
	if err := json.Unmarshal(body, &openAIError); err == nil && openAIError.Error.Message != "" {
		return nil, fmt.Errorf("Error de OpenAI: %s", openAIError.Error.Message)
	}

	return nil, fmt.Errorf("Error desconocido en la respuesta de OpenAI")
}

const openaiModelsURL = "https://api.openai.com/v1/models"

type Model struct {
	ID    string `json:"id"`
	Owned bool   `json:"owned"`
	Ready bool   `json:"ready"`
}

func (os *OpenAIService) ModelsAvailable() {
	req, err := http.NewRequest("GET", openaiModelsURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+os2.Getenv("OPEN_API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var models []Model
	if err := json.Unmarshal(body, &models); err != nil {
		fmt.Println("Error parsing response:", err)
		return
	}

	fmt.Println("Modelos disponibles:")
	for _, model := range models {
		fmt.Printf("ID: %s, Propio: %t, Listo: %t\n", model.ID, model.Owned, model.Ready)
	}
}
