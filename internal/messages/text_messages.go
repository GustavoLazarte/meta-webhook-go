package messages

type Message struct {
	MessagingProduct string `json:"messaging_product"`
	To               string `json:"to"`
	Type             string `json:"type"`
	Text             Text   `json:"text"`
}

type Text struct {
	Body string `json:"body"`
}
