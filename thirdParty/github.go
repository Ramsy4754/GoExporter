package thirdParty

type GithubRequest struct {
	ApiUrl     string `json:"apiUrl"`
	Token      string `json:"token"`
	Repository string `json:"repository"`
}

type GithubIssue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
