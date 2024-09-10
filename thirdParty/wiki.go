package thirdParty

type WikiRequest struct {
	InstanceUrl string `json:"instanceUrl"`
	UserName    string `json:"userName"`
	Token       string `json:"token"`
	SpaceKey    string `json:"spaceKey"`
}

type WikiPage struct {
	Type  string        `json:"type"`
	Title string        `json:"title"`
	Space WikiPageSpace `json:"space"`
	Body  WikiPageBody  `json:"body"`
}

type WikiPageSpace struct {
	Key string `json:"key"`
}

type WikiPageBody struct {
	Storage WikiPageBodyStorage `json:"storage"`
}

type WikiPageBodyStorage struct {
	Value          string `json:"value"`
	Representation string `json:"representation"`
}
