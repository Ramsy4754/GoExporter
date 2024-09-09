package thirdParty

type GithubRequest struct {
	Token      string `json:"token"`
	Repository string `json:"repository"`
}

type GithubIssue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
