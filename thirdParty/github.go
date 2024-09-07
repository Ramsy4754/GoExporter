package thirdParty

type GithubRequest struct {
	ApiUrl     string `json:"apiUrl"`
	Repository string `json:"repository"`
	Token      string `json:"token"`
}

type GithubIssue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
