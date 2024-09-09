package thirdParty

type GitlabRequest struct {
	ApiUrl    string `json:"apiUrl"`
	Token     string `json:"token"`
	ProjectId string `json:"projectId"`
}

type GitlabIssue struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
