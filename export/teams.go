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
	tm := new(thirdParty.TeamsMessage)

	tm.Type = "MessageCard"
	tm.Context = "http://schema.org/extensions"
	tm.ThemeColor = "0076D7"
	tm.Summary = fmt.Sprintf("CWPP Scan Result: %s", result.ScanGroupName)

	facts := make([]thirdParty.TeamsFact, 0)
	facts = append(facts, thirdParty.TeamsFact{
		Name:  "User ID",
		Value: result.UserId,
	})
	facts = append(facts, thirdParty.TeamsFact{
		Name:  "Provider",
		Value: result.Provider,
	})
	facts = append(facts, thirdParty.TeamsFact{
		Name:  "Key Name",
		Value: result.KeyName,
	})
	facts = append(facts, thirdParty.TeamsFact{
		Name:  "Event Time",
		Value: result.EventTime,
	})

	summary := fmt.Sprintf("Total: %d(%s)\nCritical: %d(%s)\nHigh: %d(%s)\nMedium: %d(%s)\nLow: %d(%s)",
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
	facts = append(facts, thirdParty.TeamsFact{
		Name:  "Summary",
		Value: summary,
	})
	section := thirdParty.TeamsSection{
		ActivityTitle: fmt.Sprintf("CWPP Scan Start: %s", result.ScanGroupName),
		Facts:         facts,
		Markdown:      true,
	}
	tm.Sections = append(tm.Sections, section)
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
	tm := new(thirdParty.TeamsMessage)

	tm.Type = "MessageCard"
	tm.Context = "http://schema.org/extensions"
	tm.ThemeColor = "0076D7"
	tm.Summary = fmt.Sprintf("CWPP Scan Start: %s", start.ScanGroupName)

	facts := make([]thirdParty.TeamsFact, 0)
	facts = append(facts, thirdParty.TeamsFact{
		Name:  "User ID",
		Value: start.UserId,
	})
	facts = append(facts, thirdParty.TeamsFact{
		Name:  "Provider",
		Value: start.Provider,
	})
	facts = append(facts, thirdParty.TeamsFact{
		Name:  "Key Name",
		Value: start.KeyName,
	})
	facts = append(facts, thirdParty.TeamsFact{
		Name:  "Event Time",
		Value: start.EventTime,
	})
	section := thirdParty.TeamsSection{
		ActivityTitle: fmt.Sprintf("CWPP Scan Start: %s", start.ScanGroupName),
		Facts:         facts,
		Markdown:      true,
	}
	tm.Sections = append(tm.Sections, section)

	return tm
}
