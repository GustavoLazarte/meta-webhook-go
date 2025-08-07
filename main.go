package main

import (
	"fmt"
	"log"
	"net/http"

	"webhook/internal/handlers"
)

// main es el punto de entrada de la aplicación.
// Solo configura y arranca el servidor HTTP, delegando toda la lógica a los paquetes internos.
func main() {
	// Token de verificación para la configuración del webhook en Meta (WhatsApp Business)
	verifyToken := "TU_TOKEN_VERIFICACION" // <-- Cambia esto por tu token real

	// Registrar el handler principal del webhook, que delega en la lógica de negocio
	http.HandleFunc("/webhook", handlers.WebhookHandler(verifyToken, handlers.BusinessMessageHandler))

	// Iniciar el servidor HTTP
	fmt.Println("Servidor escuchando en :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
