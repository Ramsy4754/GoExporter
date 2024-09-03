package thirdParty

type SlackRequest struct {
	WebhookUrl string
}

type SlackMessageText struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type SlackMessageBlock struct {
	Type string            `json:"type"`
	Text *SlackMessageText `json:"text,omitempty"`
}

type SlackMessage struct {
	Blocks []*SlackMessageBlock `json:"blocks"`
}
