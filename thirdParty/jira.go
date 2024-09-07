package thirdParty

type JiraRequest struct {
	InstanceUrl string
	ApiKey      string
	ProjectKey  string
	UserName    string
}

type JiraRequestBody struct {
	Fields JiraRequestFields `json:"fields"`
}

type JiraRequestFields struct {
	Project struct {
		Key string `json:"key"`
	} `json:"project"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Issuetype   struct {
		Name string `json:"name"`
	} `json:"issuetype"`
}
