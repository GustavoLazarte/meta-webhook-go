package messages_factory

import "webhook/internal/messages"

func BuildMessageText(to string, text string) *messages.Message {
	return &messages.Message{
		MessagingProduct: "whatsapp",
		To:               to, // Ejemplo de valor dinámico
		Type:             "text",
		Text: messages.Text{
			Body: text,
		},
	}
}

func BuildMessageInteractive(to string) *messages.InteractiveMessage {
	return &messages.InteractiveMessage{
		MessagingProduct: "whatsapp",
		To:               to,
		Type:             "interactive",
		Interactive: messages.Interactive{
			Type: "button",
			Body: messages.Body{
				Text: "Boton de Texti",
			},
			Action: messages.Action{
				Buttons: []messages.Button{
					{
						Type: "reply",
						Reply: messages.Reply{
							ID:    "buttonID1",
							Title: "buttonTitle1",
						},
					},
					{
						Type: "reply",
						Reply: messages.Reply{
							ID:    "buttonID2",
							Title: "buttonTitle2",
						},
					},
				},
			},
		},
	}
}

func BuildMessageTemplate(template_name string, recipient string, params ...string) *messages.TemplateBody {
	return &messages.TemplateBody{
		MessagingProduct: "whatsapp",
		To:               recipient,
		Type:             "template",
		Template: messages.Template{
			Name: template_name, // Nombre del template creado en Meta
			Language: messages.Language{
				Code: "en",
			},
			Components: []messages.Component{
				{
					Type: "header", // Cambia el tipo a 'header'
					Parameters: []messages.Parameter{
						{
							Type: "image", // Indica que se trata de una imagen
							Image: messages.Image{
								Link: "https://37d0138cfe2fd591be1f3bfc688fd877.serveo.net/storage/" + params[1], // URL de la imagen que quieres enviar
							},
						},
					},
				},
				{
					Type: "body", // Tipo de componente, en este caso es el cuerpo del mensaje
					Parameters: []messages.Parameter{
						{
							Type: "text",
							Text: params[0], // Parámetero dinámico (puedes cambiarlo)
						},
					},
				},
			},
		},
	}

}

func BuildMessageTemplateNoParams(template_name string, recipient string) *messages.TemplateBody {
	return &messages.TemplateBody{
		MessagingProduct: "whatsapp",
		To:               recipient,
		Type:             "template",
		Template: messages.Template{
			Name: template_name, // Nombre del template creado en Meta
			Language: messages.Language{
				Code: "en",
			},
		},
	}

}

//https://a48da63c148c1b3f37c0aed9cda55b08.serveo.net
