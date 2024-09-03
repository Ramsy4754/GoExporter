package scan

type Vulnerability struct {
	Cve         string `json:"cve"`
	Severity    string `json:"severity"`
	Description string `json:"description"`
}

type Result struct {
	ScanType        string          `json:"scanType"`
	Vulnerabilities []Vulnerability `json:"vulnerabilities"`
}
