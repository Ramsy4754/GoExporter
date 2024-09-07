package mq

type Message struct {
	Application string      `json:"application"`
	WebhookUrl  string      `json:"webhookUrl"`
	InstanceUrl string      `json:"instanceUrl"`
	ApiKey      string      `json:"apiKey"`
	ProjectKey  string      `json:"projectKey"`
	UserName    string      `json:"userName"`
	Event       string      `json:"event"`
	Args        interface{} `json:"args"`
}
