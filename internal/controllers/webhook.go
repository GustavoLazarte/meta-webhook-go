package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"webhook/internal/models"
)

type Webhook struct {
}

func (wh *Webhook) AllowAccess(w http.ResponseWriter, r *http.Request) {
	verifyToken := r.URL.Query().Get("hub.verify_token")
	challenge := r.URL.Query().Get("hub.challenge")
	fmt.Println("helllo from token", verifyToken)
	if verifyToken == os.Getenv("TOKEN_META_WH") {
		fmt.Fprint(w, challenge)
	} else {
		http.Error(w, "Token de verificación inválido", http.StatusUnauthorized)
	}
}

func (wh *Webhook) HandleWebhookEvent(w http.ResponseWriter, r *http.Request) {

	var event models.WhatsAppWebhookRequest
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		log.Printf("Error procesando el evento: %v", err)
		http.Error(w, "Error procesando el evento", http.StatusBadRequest)
		return
	}

	

	w.WriteHeader(http.StatusOK)
}
