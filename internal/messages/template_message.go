package messages

type TemplateBody struct {
	MessagingProduct string   `json:"messaging_product"`
	To               string   `json:"to"`
	Type             string   `json:"type"`
	Template         Template `json:"template"`
}

type Template struct {
	Name       string      `json:"name"`
	Language   Language    `json:"language"`
	Components []Component `json:"components"`
}

type Language struct {
	Code string `json:"code"`
}

type Component struct {
	Type       string      `json:"type"`
	Parameters []Parameter `json:"parameters"`
}

type Parameter struct {
	Type  string `json:"type"`
	Image Image  `json:"image,omitempty"` // Agregamos el tipo de imagen
	Text  string `json:"text,omitempty"`
}

type Image struct {
	Link string `json:"link"` // URL de la imagen
}
