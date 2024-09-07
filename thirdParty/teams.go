package thirdParty

type TeamsRequest struct {
	WebhookUrl string `json:"webhookUrl"`
	TenantId   string `json:"tenantId"`
}

type TeamsMessage struct {
	Type       string         `json:"@type"`
	Context    string         `json:"context"`
	ThemeColor string         `json:"themeColor"`
	Summary    string         `json:"summary"`
	Sections   []TeamsSection `json:"sections"`
}

type TeamsSection struct {
	ActivityTitle string      `json:"activityTitle"`
	Facts         []TeamsFact `json:"facts"`
	Markdown      bool        `json:"markdown"`
}

type TeamsFact struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
