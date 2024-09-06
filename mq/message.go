package mq

type Message struct {
	Application string      `json:"application"`
	WebhookUrl  string      `json:"webhookUrl"`
	Event       string      `json:"event"`
	Args        interface{} `json:"args"`
}
