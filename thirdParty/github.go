package thirdParty

type GitRequest struct {
	ApiUrl     string `json:"apiUrl"`
	Repository string `json:"repository"`
	Token      string `json:"token"`
}

type GitIssue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
