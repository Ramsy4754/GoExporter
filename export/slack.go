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

func SendMessageToSlack(request *thirdParty.SlackRequest, result *scan.Result) {
	logger := xLogger.GetLogger()

	payload := formatMessage(result)

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	_, err = http.Post(request.WebhookUrl, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		logger.Print("fail to send message to slack", err)
	}
}

func formatMessage(result *scan.Result) (sm thirdParty.SlackMessage) {
	sm.Blocks = append(sm.Blocks, &thirdParty.SlackMessageBlock{
		Type: "section",
		Text: &thirdParty.SlackMessageText{
			Type: "mrkdwn",
			Text: fmt.Sprintf("*CWPP Scan Result: %s*", result.ScanType),
		},
	})

	sm.Blocks = append(sm.Blocks, &thirdParty.SlackMessageBlock{
		Type: "divider",
	})

	for _, vuln := range result.Vulnerabilities {
		sm.Blocks = append(sm.Blocks, &thirdParty.SlackMessageBlock{
			Type: "section",
			Text: &thirdParty.SlackMessageText{
				Type: "mrkdwn",
				Text: fmt.Sprintf("*CVE:* %s\n*Severity:* %s\n*Description:* %s", vuln.Cve, vuln.Severity, vuln.Description),
			},
		})
	}
	return
}
