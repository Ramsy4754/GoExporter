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
	ProjectId   string      `json:"projectId"`
	Event       string      `json:"event"`
	Args        interface{} `json:"args"`
}

type Args struct {
	Provider      string `json:"provider"`
	UserId        string `json:"userId"`
	ScanGroupName string `json:"scanGroupName"`
	KeyName       string `json:"keyName"`
	EventTime     string `json:"eventTime"`
}

type SummaryDetail struct {
	Count      int    `json:"count"`
	Percentage string `json:"percentage"`
}

type Summary struct {
	Total    SummaryDetail `json:"total"`
	Critical SummaryDetail `json:"critical"`
	High     SummaryDetail `json:"high"`
	Medium   SummaryDetail `json:"medium"`
	Low      SummaryDetail `json:"low"`
}

type CompleteMessage struct {
	Application string  `json:"application"`
	WebhookUrl  string  `json:"webhookUrl"`
	Event       string  `json:"event"`
	Args        Args    `json:"args"`
	Summary     Summary `json:"summary"`
}
