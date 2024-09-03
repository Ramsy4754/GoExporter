package mq

import "GoExporter/scan"

type Message struct {
	Application string      `json:"application"`
	WebhookUrl  string      `json:"webhookUrl"`
	Result      scan.Result `json:"result"`
}
