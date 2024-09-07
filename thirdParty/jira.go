package thirdParty

type JiraRequest struct {
	InstanceUrl string `json:"instanceUrl"`
	ApiKey      string `json:"apiKey"`
	ProjectKey  string `json:"projectKey"`
	UserName    string `json:"userName"`
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
