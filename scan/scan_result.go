package scan

type Vulnerability struct {
	Cve         string `json:"cve"`
	Severity    string `json:"severity"`
	Description string `json:"description"`
}

type ResultInfo struct {
	Provider      string `json:"provider"`
	UserId        string `json:"userId"`
	ScanGroupName string `json:"scanGroupName"`
	KeyName       string `json:"keyName"`
	EventTime     string `json:"eventTIme"`
	ResultSummary `json:"summary"`
}

type ResultSummary struct {
	Total    ResultSummaryData `json:"total"`
	Critical ResultSummaryData `json:"critical"`
	High     ResultSummaryData `json:"high"`
	Medium   ResultSummaryData `json:"medium"`
	Low      ResultSummaryData `json:"low"`
}

type ResultSummaryData struct {
	Count      int    `json:"count"`
	Percentage string `json:"percentage"`
}

type StartInfo struct {
	Provider      string `json:"provider"`
	UserId        string `json:"userId"`
	ScanGroupName string `json:"scanGroupName"`
	KeyName       string `json:"keyName"`
	EventTime     string `json:"eventTIme"`
}
