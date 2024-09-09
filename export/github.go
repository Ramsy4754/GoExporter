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

func SendCwppScanStartToGithub(request *thirdParty.GithubRequest, start *scan.StartInfo) {
	issue := formatCwppScanStartGithubIssue(start)
	sendToGithub(request, issue)
}

func sendToGithub(request *thirdParty.GithubRequest, issue thirdParty.GithubIssue) {
	logger := xLogger.GetLogger()

	payloadBytes, err := json.Marshal(issue)
	if err != nil {
		logger.Printf("failed to marshal issue: %v", err)
		return
	}

	requestUrl := fmt.Sprintf("%s/projects/%s/issues", request.ApiUrl, request.Repository)
	req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(payloadBytes))
	if err != nil {
		logger.Print("failed to create request: ", err)
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("token %s", request.Token))

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

func formatCwppScanStartGithubIssue(start *scan.StartInfo) thirdParty.GithubIssue {
	comment := fmt.Sprintf("## CWPP Scan Start: %s\n\n### User ID: %s\n**Provider:** %s\n**Key Name:** %s\n**Event Time**: %s\n\n",
		start.ScanGroupName,
		start.UserId,
		start.Provider,
		start.KeyName,
		start.EventTime,
	)

	issue := thirdParty.GithubIssue{
		Title: fmt.Sprintf("CWPP Scan Start: %s", start.ScanGroupName),
		Body:  comment,
	}
	return issue
}

func SendCwppScanResultToGithub(request *thirdParty.GithubRequest, result *scan.ResultInfo) {
	issue := formatCwppScanResultGithubIssue(result)
	sendToGithub(request, issue)
}

func formatCwppScanResultGithubIssue(result *scan.ResultInfo) thirdParty.GithubIssue {
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

	issue := thirdParty.GithubIssue{
		Title: fmt.Sprintf("CWPP Scan Result: %s", result.ScanGroupName),
		Body:  comment,
	}
	return issue
}
