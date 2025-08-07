package models

type WhatsaapClient struct {
	ClientID    string `bson:"client_id"`    // ID del usuario
	State       string `bson:"state"`        // Estado del client actual
	CurrentFlow string `bson:"current_flow"` // Estado del chat actual
}
