package mq

type Message struct {
	Application string      `json:"application"`
	WebhookUrl  string      `json:"webhookUrl"`
	TenantId    string      `json:"tenantId"`
	InstanceUrl string      `json:"instanceUrl"`
	ApiKey      string      `json:"apiKey"`
	ProjectKey  string      `json:"projectKey"`
	UserName    string      `json:"userName"`
	Repository  string      `json:"repository"`
	Token       string      `json:"token"`
	SpaceKey    string      `json:"spaceKey"`
	ApiUrl      string      `json:"apiUrl"`
	ProjectId   string      `json:"projectId"`
	Event       string      `json:"event"`
	Args        interface{} `json:"args"`
}
