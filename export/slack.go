package export

import (
	"GoExporter/scan"
	"GoExporter/thirdParty"
	"GoExporter/xLogger"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendCwppScanResultToSlack(request *thirdParty.SlackRequest, result *scan.Result) {
	logger := xLogger.GetLogger()

	payload := formatCwppScanResultMessage(result)

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	resp, err := http.Post(request.WebhookUrl, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		logger.Print("fail to send message to slack", err)
	}
	logger.Print("response status:", resp.Status)
}

func formatCwppScanResultMessage(result *scan.Result) (sm thirdParty.SlackMessage) {
	//sm.Blocks = append(sm.Blocks, &thirdParty.SlackMessageBlock{
	//	Type: "section",
	//	Text: &thirdParty.SlackMessageText{
	//		Type: "mrkdwn",
	//		Text: fmt.Sprintf("*CWPP Scan Result: %s*", result.ScanType),
	//	},
	//})
	//
	//sm.Blocks = append(sm.Blocks, &thirdParty.SlackMessageBlock{
	//	Type: "divider",
	//})
	//
	//for _, vuln := range result.Vulnerabilities {
	//	sm.Blocks = append(sm.Blocks, &thirdParty.SlackMessageBlock{
	//		Type: "section",
	//		Text: &thirdParty.SlackMessageText{
	//			Type: "mrkdwn",
	//			Text: fmt.Sprintf("*CVE:* %s\n*Severity:* %s\n*Description:* %s", vuln.Cve, vuln.Severity, vuln.Description),
	//		},
	//	})
	//}
	return
}

func SendCwppScanStartToSlack(request *thirdParty.SlackRequest, start *scan.StartInfo) {
	logger := xLogger.GetLogger()

	payload := formatCwppScanStartMessage(start)

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		logger.Println("error marshalling json:", err)
		return
	}

	resp, err := http.Post(request.WebhookUrl, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		logger.Print("failed to send message to slack", err)
		return
	}
	logger.Print("response status:", resp.Status)
}

func formatCwppScanStartMessage(start *scan.StartInfo) (sm thirdParty.SlackMessage) {
	sm.Blocks = append(sm.Blocks, &thirdParty.SlackMessageBlock{
		Type: "section",
		Text: &thirdParty.SlackMessageText{
			Type: "mrkdwn",
			Text: fmt.Sprintf("*CWPP Scan Start*:  %s\n", start.ScanGroupName),
		},
	})

	sm.Blocks = append(sm.Blocks, &thirdParty.SlackMessageBlock{
		Type: "divider",
	})

	sm.Blocks = append(sm.Blocks, &thirdParty.SlackMessageBlock{
		Type: "section",
		Text: &thirdParty.SlackMessageText{
			Type: "mrkdwn",
			Text: fmt.Sprintf("*Provider*: %s\n*User ID*: %s\n*Key Name*: %s\n", start.Provider, start.UserId, start.KeyName),
		},
	})

	return
}
