package requests

type WhatsappClient struct {
	ClientID    string `json:"client_id"`    // ID del usuario
	State       string `json:"state"`        // Estado del client actual
	CurrentFlow string `json:"current_flow"` // Estado del chat actual
}
