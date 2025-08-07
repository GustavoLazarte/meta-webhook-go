package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webhook/internal/models"
)
// SendWhatsAppMessage envía un mensaje de texto a través de la API de WhatsApp (Meta)
func SendWhatsAppMessage(phoneNumberID, accessToken, to, message string) error {
	url := fmt.Sprintf("https://graph.facebook.com/v19.0/%s/messages", phoneNumberID)
	body := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to": to,
		"type": "text",
		"text": map[string]string{"body": message},
	}
	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return fmt.Errorf("error enviando mensaje: %s", resp.Status)
	}
	return nil
}

// WhatsAppWebhookRequest ahora es un alias del modelo
type WhatsAppWebhookRequest = models.WhatsAppWebhookRequest

// MessageHandler define la firma para manejar mensajes entrantes
type MessageHandler func(userID, msgType string, payload models.WhatsAppWebhookRequest) error

// HandleIncomingMessage decodifica el request y llama al handler de negocio
func HandleIncomingMessage(w http.ResponseWriter, r *http.Request, handler MessageHandler) {
	var req models.WhatsAppWebhookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for _, entry := range req.Entry {
		for _, change := range entry.Changes {
			for _, msg := range change.Value.Messages {
				_ = handler(msg.From, msg.Type, req)
			}
		}
	}
	w.WriteHeader(http.StatusOK)
}
