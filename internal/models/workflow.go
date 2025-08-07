package models

// Step representa un paso de un workflow.
type Step struct {
	Name        string
	Description string
	// Action se define en la capa de dominio, no en el modelo
}

// Workflow representa la estructura de un flujo de trabajo.
type Workflow struct {
	ID      string
	Name    string
	Steps   []Step
}
