package domain

// Recipient The recipient of our message
type Recipient struct {
	ID string `json:"id"`
}

type QuickReply struct {
	ContentType string `json:"content_type,omitempty"`
	Title       string `json:"title,omitempty"`
	Payload     string `json:"payload,omitempty"`
	ImageUrl    string `json:"image_url,omitempty"`
}

// Message The message to send it its basic
type Message struct {
	Text         string       `json:"text,omitempty"`
	Mid          string       `json:"mid,omitempty"`
	QuickReplies []QuickReply `json:"quick_replies,omitempty"`
}

type Button struct {
	Type    string `json:"type,omitempty"`
	Title   string `json:"title,omitempty"`
	Payload string `json:"payload,omitempty"`
	URL     string `json:"url,omitempty"`
}

type Element struct {
	Title         string        `json:"title,omitempty"`
	Subtitle      string        `json:"subtitle,omitempty"`
	ImageURL      string        `json:"image_url,omitempty"`
	DefaultAction DefaultAction `json:"default_action,omitempty"`
	Buttons       []Button      `json:"buttons,omitempty"`
}

type DefaultAction struct {
	Type                string `json:"type,omitempty"`
	URL                 string `json:"url,omitempty"`
	WebViewHeightRation string `json:"webview_height_ratio,omitempty"`
}

// Attachment The attachment to send (custom)
type Attachment struct {
	Attachment struct {
		Type    string `json:"type,omitempty"`
		Payload struct {
			TemplateType string    `json:"template_type,omitempty"`
			Elements     []Element `json:"elements,omitempty"`
			URL          string    `json:"url,omitempty"`
		} `json:"payload,omitempty"`
	} `json:"attachment,omitempty"`
}

// ResponseAttachment Full response
type ResponseAttachment struct {
	Recipient   Recipient  `json:"recipient"`
	MessageType string     `json:"message_type,omitempty"`
	Message     Attachment `json:"message,omitempty"`
}

// ResponseMessage Full response
type ResponseMessage struct {
	Recipient     Recipient `json:"recipient"`
	MessagingType string    `json:"messaging_type,omitempty"`
	Message       Message   `json:"message,omitempty"`
}
