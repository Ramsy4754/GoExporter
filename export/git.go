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

func SendCwppScanStartToGit(request *thirdParty.GitRequest, start *scan.StartInfo) {
	issue := formatCwppScanStartGithubIssue(start)
	sendToGithub(request, issue)
}

func sendToGithub(request *thirdParty.GitRequest, issue thirdParty.GitIssue) {
	payloadBytes, err := json.Marshal(issue)
	logger := xLogger.GetLogger()
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

func formatCwppScanStartGithubIssue(start *scan.StartInfo) thirdParty.GitIssue {
	comment := fmt.Sprintf("## CWPP Scan Start: %s\n\n### User ID: %s\n**Provider:** %s\n**Scan Group Name:** %s\n**Key Name:** %s\n**Key Value:** %s\n\n",
		start.ScanGroupName,
		start.UserId,
		start.Provider,
		start.ScanGroupName,
		start.KeyName,
	)

	issue := thirdParty.GitIssue{
		Title: fmt.Sprintf("CWPP Scan Start: %s", start.ScanGroupName),
		Body:  comment,
	}
	return issue
}

func SendCwppScanResultToGit(request *thirdParty.GitRequest, result *scan.ResultInfo) {
	issue := formatCwppScanResultGithubIssue(result)
	sendToGithub(request, issue)
}

func formatCwppScanResultGithubIssue(result *scan.ResultInfo) thirdParty.GitIssue {
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
			"**Scan Group Name:** %s\n"+
			"**Key Name:** %s\n"+
			"**Event Time:** %s\n\n%s",
		result.ScanGroupName,
		result.UserId,
		result.Provider,
		result.ScanGroupName,
		result.KeyName,
		result.EventTime,
		summary,
	)

	issue := thirdParty.GitIssue{
		Title: fmt.Sprintf("CWPP Scan Result: %s", result.ScanGroupName),
		Body:  comment,
	}
	return issue
}
