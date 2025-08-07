package handlers

import (
	"context"
	"fmt"
	"os"
	"webhook/internal/models"
	"webhook/internal/whatsapp"
	"webhook/internal/workflow"
)

// userWorkflows almacena los workflows activos por usuario (usar almacenamiento persistente en producción)
var userWorkflows = make(map[string]*workflow.Workflow)

// BusinessMessageHandler implementa la lógica de negocio para cada mensaje entrante de WhatsApp
func BusinessMessageHandler(userID, msgType string, payload models.WhatsAppWebhookRequest) error {
	wf, ok := userWorkflows[userID]
	if !ok {
		wf = workflow.New("demo", "Demo Workflow", []workflow.Step{
			{
				Step: models.Step{
					Name:        "Selecciona una opción",
					Description: "Elige una opción para continuar",
				},
				Action: func(ctx context.Context, data map[string]interface{}) error {
					return sendStepMessage(userID, "Por favor selecciona una opción: 1) Aceptar 2) Rechazar")
				},
			},
			{
				Step: models.Step{
					Name:        "Procesar selección",
					Description: "Procesa la respuesta del usuario",
				},
				Action: func(ctx context.Context, data map[string]interface{}) error {
					if sel, ok := data["seleccion"].(string); ok {
						return sendStepMessage(userID, "Recibido: "+sel+". ¡Gracias por tu respuesta!")
					}
					return sendStepMessage(userID, "No se recibió selección válida.")
				},
			},
		})
		userWorkflows[userID] = wf
	}

	if msgType == "interactive" && len(payload.Entry) > 0 && len(payload.Entry[0].Changes) > 0 && len(payload.Entry[0].Changes[0].Value.Messages) > 0 {
		msg := payload.Entry[0].Changes[0].Value.Messages[0]
		wf.State["seleccion"] = msg.Interactive.ListReply.Title
		_ = wf.NextStep(context.Background())
	} else if msgType == "text" {
		wf.Reset()
		_ = wf.NextStep(context.Background())
	}
	return nil
}

// sendStepMessage envía un mensaje de texto a WhatsApp usando variables de entorno para credenciales
func sendStepMessage(to, message string) error {
	phoneNumberID := os.Getenv("WA_PHONE_NUMBER_ID")
	accessToken := os.Getenv("WA_ACCESS_TOKEN")
	if phoneNumberID == "" || accessToken == "" {
		fmt.Println("[WARN] Variables de entorno WA_PHONE_NUMBER_ID o WA_ACCESS_TOKEN no definidas")
		return nil
	}
	return whatsapp.SendWhatsAppMessage(phoneNumberID, accessToken, to, message)
}
