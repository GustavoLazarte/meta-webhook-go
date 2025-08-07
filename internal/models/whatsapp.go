package models

// WhatsAppWebhookRequest representa la estructura básica de un webhook de WhatsApp (Meta API)
type WhatsAppWebhookRequest struct {
	Entry []Entry `json:"entry"`
}

type Entry struct {
	Changes []Change `json:"changes"`
}

type Change struct {
	Value Value `json:"value"`
}

type Value struct {
	Messages []WhatsAppMessage `json:"messages"`
}

type WhatsAppMessage struct {
	From        string      `json:"from"`
	ID          string      `json:"id"`
	Type        string      `json:"type"`
	Text        Text        `json:"text"`
	Interactive Interactive `json:"interactive"`
}

type Text struct {
	Body string `json:"body"`
}

type Interactive struct {
	Type      string    `json:"type"`
	ListReply ListReply `json:"list_reply"`
}

type ListReply struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
