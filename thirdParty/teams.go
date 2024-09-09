package thirdParty

type TeamsRequest struct {
	WebhookUrl string `json:"webhookUrl"`
	TenantId   string `json:"tenantId"`
}

type TeamsMessage struct {
	Type        string            `json:"type"`
	Attachments []TeamsAttachment `json:"attachments"`
}

type TeamsAttachment struct {
	ContentType string       `json:"contentType"`
	Content     TeamsContent `json:"content"`
}

type TeamsContent struct {
	Type    string             `json:"type"`
	Version string             `json:"version"`
	Schema  string             `json:"@schema"`
	Body    []TeamsContentBody `json:"body"`
	Actions []string           `json:"actions"`
}

type TeamsContentBody struct {
	Type   string      `json:"type"`
	Size   string      `json:"size"`
	Weight string      `json:"weight"`
	Text   string      `json:"text"`
	Facts  []TeamsFact `json:"facts"`
}

type TeamsFact struct {
	Title string `json:"title"`
	Value string `json:"value"`
}
