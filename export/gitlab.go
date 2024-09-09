package export

import (
	"GoExporter/scan"
	"GoExporter/thirdParty"
	"GoExporter/xLogger"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SendCwppScanResultToGitlab(request *thirdParty.GitlabRequest, result *scan.ResultInfo) {
	issue := formatCwppScanResultGitlabIssue(result)
	sendToGitlab(request, issue)
}

func formatCwppScanResultGitlabIssue(result *scan.ResultInfo) (gi thirdParty.GitlabIssue) {
	summary := fmt.Sprintf(
		"### Result Summary\n**Total:** %d\n**Critical:** %d\n**High:** %d\n**Medium:** %d\n**Low:** %d\n\n",
		result.Total.Count,
		result.Critical.Count,
		result.High.Count,
		result.Medium.Count,
		result.Low.Count,
	)

	comment := fmt.Sprintf(
		"## CWPP Scan Result: %s\n\n"+
			"### User ID: %s\n"+
			"**Provider:** %s\n"+
			"**Key Name:** %s\n"+
			"**Event Time:** %s\n\n%s",
		result.ScanGroupName,
		result.UserId,
		result.Provider,
		result.KeyName,
		result.EventTime,
		summary,
	)

	gi.Title = fmt.Sprintf("CWPP Scan Result: %s", result.ScanGroupName)
	gi.Description = comment

	return
}

func SendCwppScanStartToGitlab(request *thirdParty.GitlabRequest, start *scan.StartInfo) {
	issue := formatCwppScanStartGitlabIssue(start)
	sendToGitlab(request, issue)
}

func sendToGitlab(request *thirdParty.GitlabRequest, issue thirdParty.GitlabIssue) {
	logger := xLogger.GetLogger()

	payloadBytes, err := json.Marshal(issue)
	if err != nil {
		logger.Printf("failed to marshal issue: %v", err)
		return
	}

	requestUrl := fmt.Sprintf("https://gitlab.com/api/v4/projects/%s/issues", request.ProjectId)
	req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(payloadBytes))
	if err != nil {
		logger.Print("failed to create request: ", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("PRIVATE-TOKEN", request.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Print("failed to send request: ", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logger.Print("failed to close response body: ", err)
		}
	}(resp.Body)

	logger.Printf("response Status: %s", resp.Status)
}

func formatCwppScanStartGitlabIssue(start *scan.StartInfo) (gi thirdParty.GitlabIssue) {
	gi.Title = fmt.Sprintf("CWPP Scan Start: %s", start.ScanGroupName)
	comment := fmt.Sprintf("## CWPP Scan Start: %s\n\n### User ID: %s\n**Provider:** %s\n**Key Name:** %s\n**Event Time**: %s\n\n",
		start.ScanGroupName,
		start.UserId,
		start.Provider,
		start.KeyName,
		start.EventTime,
	)
	gi.Description = comment
	return
}
