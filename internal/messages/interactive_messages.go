package messages

// Structs para el JSON proporcionado
type InteractiveMessage struct {
	MessagingProduct string      `json:"messaging_product"`
	To               string      `json:"to"`
	Type             string      `json:"type"`
	Interactive      Interactive `json:"interactive"`
}

type Interactive struct {
	Type   string `json:"type"`
	Body   Body   `json:"body"`
	Action Action `json:"action"`
}

type Body struct {
	Text string `json:"text"`
}

type Action struct {
	Buttons []Button `json:"buttons"`
}

type Button struct {
	Type  string `json:"type"`
	Reply Reply  `json:"reply"`
}

type Reply struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
