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

func SendCwppScanResultToTeams(request *thirdParty.TeamsRequest, result *scan.ResultInfo) {
	payload := formatCwppScanResultTeamsMessage(result)
	sendToTeams(request, payload)
}

func formatCwppScanResultTeamsMessage(result *scan.ResultInfo) *thirdParty.TeamsMessage {
	tm := &thirdParty.TeamsMessage{
		Type:        "message",
		Attachments: make([]thirdParty.TeamsAttachment, 1),
	}

	tm.Attachments[0].ContentType = "application/vnd.microsoft.card.adaptive"
	tm.Attachments[0].Content = thirdParty.TeamsContent{
		Type:    "AdaptiveCard",
		Version: "1.2",
		Schema:  "http://adaptivecards.io/schemas/adaptive-card.json",
		Body:    []thirdParty.TeamsContentBody{},
	}

	body := &tm.Attachments[0].Content.Body
	*body = append(*body, thirdParty.TeamsContentBody{
		Type:   "TextBlock",
		Size:   "Medium",
		Weight: "Bolder",
		Text:   fmt.Sprintf("CWPP Scan Result: %s", result.ScanGroupName),
	})
	*body = append(*body, thirdParty.TeamsContentBody{
		Type: "FactSet",
		Facts: []thirdParty.TeamsFact{
			{
				Title: "Scan Group Name",
				Value: result.ScanGroupName,
			},
		},
	})

	summary := fmt.Sprintf(
		"Result Summary\nTotal: %d\nCritical: %d\nHigh: %d\nMedium: %d\nLow: %d\n\n",
		result.Total.Count,
		result.Critical.Count,
		result.High.Count,
		result.Medium.Count,
		result.Low.Count,
	)

	*body = append(*body, thirdParty.TeamsContentBody{
		Type: "FactSet",
		Facts: []thirdParty.TeamsFact{
			{
				Title: "User ID",
				Value: result.UserId,
			},
			{
				Title: "Provider",
				Value: result.Provider,
			},
			{
				Title: "Key Name",
				Value: result.KeyName,
			},
			{
				Title: "Event Time",
				Value: result.EventTime,
			},
			{
				Title: "Summary",
				Value: summary,
			},
		},
	})

	return tm
}

func SendCwppScanStartToTeams(request *thirdParty.TeamsRequest, start *scan.StartInfo) {
	payload := formatCwppScanStartTeamsMessage(start)
	sendToTeams(request, payload)
}

func sendToTeams(request *thirdParty.TeamsRequest, payload *thirdParty.TeamsMessage) {
	logger := xLogger.GetLogger()

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		logger.Println("failed to marshal teams message:", err)
		return
	}

	resp, err := http.Post(request.WebhookUrl, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		logger.Println("failed to send teams message:", err)
	}
	logger.Println("teams message sent:", resp)
}

func formatCwppScanStartTeamsMessage(start *scan.StartInfo) *thirdParty.TeamsMessage {
	tm := &thirdParty.TeamsMessage{
		Type:        "message",
		Attachments: make([]thirdParty.TeamsAttachment, 1),
	}

	tm.Attachments[0].ContentType = "application/vnd.microsoft.card.adaptive"
	tm.Attachments[0].Content = thirdParty.TeamsContent{
		Type:    "AdaptiveCard",
		Version: "1.2",
		Schema:  "http://adaptivecards.io/schemas/adaptive-card.json",
		Body:    []thirdParty.TeamsContentBody{},
	}

	body := &tm.Attachments[0].Content.Body
	*body = append(*body, thirdParty.TeamsContentBody{
		Type:   "TextBlock",
		Size:   "Medium",
		Weight: "Bolder",
		Text:   fmt.Sprintf("CWPP Scan Start: %s", start.ScanGroupName),
	})
	*body = append(*body, thirdParty.TeamsContentBody{
		Type: "FactSet",
		Facts: []thirdParty.TeamsFact{
			{
				Title: "Scan Group Name",
				Value: start.ScanGroupName,
			},
		},
	})
	*body = append(*body, thirdParty.TeamsContentBody{
		Type: "FactSet",
		Facts: []thirdParty.TeamsFact{
			{
				Title: "User ID",
				Value: start.UserId,
			},
			{
				Title: "Provider",
				Value: start.Provider,
			},
			{
				Title: "Key Name",
				Value: start.KeyName,
			},
			{
				Title: "Event Time",
				Value: start.EventTime,
			},
		},
	})

	return tm
}
