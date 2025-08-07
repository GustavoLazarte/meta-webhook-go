package handlers

import (
	"net/http"
	"webhook/internal/whatsapp"
)

// WebhookHandler maneja la verificación y mensajes entrantes de WhatsApp
func WebhookHandler(verifyToken string, msgHandler whatsapp.MessageHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			mode := r.URL.Query().Get("hub.mode")
			token := r.URL.Query().Get("hub.verify_token")
			challenge := r.URL.Query().Get("hub.challenge")
			if mode == "subscribe" && token == verifyToken {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(challenge))
				return
			}
			w.WriteHeader(http.StatusForbidden)
			return
		case http.MethodPost:
			whatsapp.HandleIncomingMessage(w, r, msgHandler)
			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}
