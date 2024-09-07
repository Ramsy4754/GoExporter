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

func SendCwppScanResultToSlack(request *thirdParty.SlackRequest, result *scan.ResultInfo) {
	payload := formatCwppScanResultSlackMessage(result)
	sendToSlack(request, payload)
}

func sendToSlack(request *thirdParty.SlackRequest, payload thirdParty.SlackMessage) {
	logger := xLogger.GetLogger()

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		logger.Println("Error marshalling JSON:", err)
		return
	}

	resp, err := http.Post(request.WebhookUrl, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		logger.Print("fail to send message to slack", err)
	}
	logger.Print("response status:", resp.Status)
}

func formatCwppScanResultSlackMessage(result *scan.ResultInfo) (sm thirdParty.SlackMessage) {
	sm.Blocks = append(sm.Blocks, &thirdParty.SlackMessageBlock{
		Type: "section",
		Text: &thirdParty.SlackMessageText{
			Type: "mrkdwn",
			Text: fmt.Sprintf("*CWPP Scan Result*:  %s\n", result.ScanGroupName),
		},
	})
	sm.Blocks = append(sm.Blocks, &thirdParty.SlackMessageBlock{
		Type: "section",
		Text: &thirdParty.SlackMessageText{
			Type: "mrkdwn",
			Text: fmt.Sprintf("*Event Time(UTC)*:  %s\n", result.EventTime),
		},
	})

	sm.Blocks = append(sm.Blocks, &thirdParty.SlackMessageBlock{
		Type: "divider",
	})

	sm.Blocks = append(sm.Blocks, &thirdParty.SlackMessageBlock{
		Type: "section",
		Text: &thirdParty.SlackMessageText{
			Type: "mrkdwn",
			Text: fmt.Sprintf("*Provider*: %s\n*User ID*: %s\n*Key Name*: %s\n", result.Provider, result.UserId, result.KeyName),
		},
	})

	sm.Blocks = append(sm.Blocks, &thirdParty.SlackMessageBlock{
		Type: "divider",
	})

	sm.Blocks = append(sm.Blocks, &thirdParty.SlackMessageBlock{
		Type: "section",
		Text: &thirdParty.SlackMessageText{
			Type: "mrkdwn",
			Text: fmt.Sprintf(
				"*Total*: %d(%s)\n*Critical*: %d(%s)\n*High*: %d(%s)\n*Medium*: %d(%s)\n*Low*: %d(%s)\n",
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
		},
	})

	return
}

func SendCwppScanStartToSlack(request *thirdParty.SlackRequest, start *scan.StartInfo) {
	payload := formatCwppScanStartSlackMessage(start)
	sendToSlack(request, payload)
}

func formatCwppScanStartSlackMessage(start *scan.StartInfo) (sm thirdParty.SlackMessage) {
	sm.Blocks = append(sm.Blocks, &thirdParty.SlackMessageBlock{
		Type: "section",
		Text: &thirdParty.SlackMessageText{
			Type: "mrkdwn",
			Text: fmt.Sprintf("*CWPP Scan Start*:  %s\n", start.ScanGroupName),
		},
	})
	sm.Blocks = append(sm.Blocks, &thirdParty.SlackMessageBlock{
		Type: "section",
		Text: &thirdParty.SlackMessageText{
			Type: "mrkdwn",
			Text: fmt.Sprintf("*Event Time(UTC)*:  %s\n", start.EventTime),
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
