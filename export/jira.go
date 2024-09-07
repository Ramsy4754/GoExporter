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

func sendToJira(request *thirdParty.JiraRequest, payload thirdParty.JiraRequestBody) {
	logger := xLogger.GetLogger()
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		logger.Println("failed to marshal payload:", err)
		return
	}

	requestUrl := fmt.Sprintf("%s/rest/api/2/issue", request.InstanceUrl)
	req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(payloadBytes))
	if err != nil {
		logger.Print("failed to create request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(request.UserName, request.ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Print("failed to send request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Print("failed to close response body:", err)
			return
		}
	}(resp.Body)

	logger.Print("response status:", resp.Status)
}

func SendCwppScanStartToJira(request *thirdParty.JiraRequest, start *scan.StartInfo) {
	payload := formatCwppScanStartJiraMessage(request, start)
	sendToJira(request, payload)
}

func formatCwppScanStartJiraMessage(request *thirdParty.JiraRequest, start *scan.StartInfo) thirdParty.JiraRequestBody {
	payload := thirdParty.JiraRequestBody{
		Fields: thirdParty.JiraRequestFields{
			Project: struct {
				Key string `json:"key"`
			}{
				Key: request.ProjectKey,
			},
			Summary: "CWPP Scan Start",
			Description: fmt.Sprintf("Provider: %s\nUser ID: %s\nScan Group Name: %s\nKey Name: %s\nEvent Time: %s",
				start.Provider,
				start.UserId,
				start.ScanGroupName,
				start.KeyName,
				start.EventTime,
			),
			Issuetype: struct {
				Name string `json:"name"`
			}{
				"Bug",
			},
		},
	}
	return payload
}

func SendCwppScanResultToJira(request *thirdParty.JiraRequest, result *scan.ResultInfo) {
	payload := formatCwppScanResultJiraMessage(request, result)
	sendToJira(request, payload)
}

func formatCwppScanResultJiraMessage(request *thirdParty.JiraRequest, result *scan.ResultInfo) thirdParty.JiraRequestBody {
	payload := thirdParty.JiraRequestBody{
		Fields: thirdParty.JiraRequestFields{
			Project: struct {
				Key string `json:"key"`
			}{
				Key: request.ProjectKey,
			},
			Summary: "CWPP Scan Result",
			Description: fmt.Sprintf("Provider: %s\nUser ID: %s\nScan Group Name: %s\nKey Name: %s\nEvent Time: %s\n\nSummary\nTotal: %d(%s)\nCritical: %d(%s)\nHigh: %d(%s)\nMedium: %d(%s)\nLow: %d(%s)\n",
				result.Provider,
				result.UserId,
				result.ScanGroupName,
				result.KeyName,
				result.EventTime,
				result.Total.Count,
				result.Total.Percentage,
				result.Critical.Count,
				result.Critical.Percentage,
				result.High.Count,
				result.High.Percentage,
				result.Medium.Count,
				result.Medium.Percentage,
				result.Low.Count,
				result.Low.Percentage,
			),
			Issuetype: struct {
				Name string `json:"name"`
			}{
				"Bug",
			},
		},
	}
	return payload
}
