package export

import (
	"GoExporter/scan"
	"GoExporter/thirdParty"
	"GoExporter/xLogger"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SendCwppScanStartToWiki(request *thirdParty.WikiRequest, start *scan.StartInfo) {
	payload := formatCwppScanStartWikiPage(request, start)
	sendToWiki(request, payload)
}

func sendToWiki(request *thirdParty.WikiRequest, payload *thirdParty.WikiPage) {
	logger := xLogger.GetLogger()
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		logger.Println("failed to marshal wiki payload:", err)
		return
	}

	req, err := http.NewRequest("POST", request.InstanceUrl, bytes.NewBuffer(payloadBytes))
	if err != nil {
		logger.Println("failed to create wiki request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	auth := request.UserName + ":" + request.ApiKey
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Authorization", "Basic "+encodedAuth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Println("failed to send wiki request:", err)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logger.Println("failed to close wiki response body:", err)
		}
	}(resp.Body)

	logger.Println("wiki response status:", resp.Status)
}

func formatCwppScanStartWikiPage(request *thirdParty.WikiRequest, start *scan.StartInfo) (wp *thirdParty.WikiPage) {
	content := fmt.Sprintf("h1. CWPP Scan Start: %s\n\nh2. User ID: %s\n*Provider: %s\n*Key Name: %s\n*Event Time: %s\n\n",
		start.ScanGroupName,
		start.UserId,
		start.Provider,
		start.KeyName,
		start.EventTime,
	)

	wp.Type = "page"
	wp.Title = fmt.Sprintf("CWPP Scan Start: %s", start.ScanGroupName)
	wp.Space = thirdParty.WikiPageSpace{Key: request.SpaceKey}
	wp.Body = thirdParty.WikiPageBody{
		Storage: thirdParty.WikiPageBodyStorage{
			Value:          content,
			Representation: "wiki",
		},
	}
	return
}

func SendCwppScanResultToWiki(request *thirdParty.WikiRequest, result *scan.ResultInfo) {
	payload := formatCwppScanResultWikiPage(request, result)
	sendToWiki(request, payload)
}

func formatCwppScanResultWikiPage(request *thirdParty.WikiRequest, result *scan.ResultInfo) (wp *thirdParty.WikiPage) {
	summary := fmt.Sprintf("h2. Summary\n*Total: %d(%s)\nCritical: %d(%s)\nHigh: %d(%s)\nMedium: %d(%s)\nLow: %d(%s)\n",
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
	)

	content := fmt.Sprintf("h1. CWPP Scan Start: %s\n\nh2. User ID: %s\n*Provider: %s\n*Key Name: %s\n*Event Time: %s\n\n%s",
		result.ScanGroupName,
		result.UserId,
		result.Provider,
		result.KeyName,
		result.EventTime,
		summary,
	)

	wp.Type = "page"
	wp.Title = fmt.Sprintf("CWPP Scan Start: %s", result.ScanGroupName)
	wp.Space = thirdParty.WikiPageSpace{Key: request.SpaceKey}
	wp.Body = thirdParty.WikiPageBody{
		Storage: thirdParty.WikiPageBodyStorage{
			Value:          content,
			Representation: "wiki",
		},
	}
	return
}
