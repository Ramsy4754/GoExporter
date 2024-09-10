package mq

import (
	"GoExporter/export"
	"GoExporter/scan"
	"GoExporter/thirdParty"
	"GoExporter/xLogger"
	"encoding/json"
	"github.com/streadway/amqp"
)

func getRabbitMQUrl() string {
	return "amqp://guest:guest@localhost:5672/"
}

func Listen() {
	logger := xLogger.GetLogger()

	url := getRabbitMQUrl()
	conn, err := amqp.Dial(url)
	if err != nil {
		logger.Fatalf("failed to connect to %s: %s", url, err)
	}

	logger.Printf("connect to %s", url)

	defer func(conn *amqp.Connection) {
		err = conn.Close()
		if err != nil {
			logger.Fatalf("failed to close connection: %s", err)
		}
	}(conn)

	ch, err := conn.Channel()
	if err != nil {
		logger.Fatalf("failed to open a channel: %s", err)
	}

	defer func(ch *amqp.Channel) {
		err = ch.Close()
		if err != nil {
			logger.Fatalf("failed to close a channel: %s", err)
		}
	}(ch)

	q, err := ch.QueueDeclare(
		"task_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Fatalf("failed to declare a queue: %s", err)
	}

	msgList, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Fatalf("failed to register a consumer: %s", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgList {
			var msg Message
			if err := json.Unmarshal(d.Body, &msg); err != nil {
				logger.Printf("error parsing JSON: %s", err)
				continue
			}

			logger.Printf("received message \n%+v", msg)

			switch msg.Application {
			case "slack":
				request := &thirdParty.SlackRequest{
					WebhookUrl: msg.WebhookUrl,
				}
				switch msg.Event {
				case "afterCwppScan":
					result := getScanResultArgsFromMsg(msg)
					export.SendCwppScanResultToSlack(request, result)
					break
				case "beforeCwppScan":
					start := getScanStartArgsFromMsg(msg)
					export.SendCwppScanStartToSlack(request, start)
				default:
					logger.Printf("unknown event: %s", msg.Event)
				}
				break
			case "jira":
				request := &thirdParty.JiraRequest{
					InstanceUrl: msg.InstanceUrl,
					ApiKey:      msg.ApiKey,
					ProjectKey:  msg.ProjectKey,
					UserName:    msg.UserName,
				}
				switch msg.Event {
				case "afterCwppScan":
					result := getScanResultArgsFromMsg(msg)
					export.SendCwppScanResultToJira(request, result)
					break
				case "beforeCwppScan":
					start := getScanStartArgsFromMsg(msg)
					export.SendCwppScanStartToJira(request, start)
					break
				default:
					logger.Printf("unknown event: %s", msg.Event)
				}
				break
			case "teams":
				request := &thirdParty.TeamsRequest{
					WebhookUrl: msg.WebhookUrl,
					TenantId:   msg.TenantId,
				}
				switch msg.Event {
				case "afterCwppScan":
					result := getScanResultArgsFromMsg(msg)
					export.SendCwppScanResultToTeams(request, result)
					break
				case "beforeCwppScan":
					start := getScanStartArgsFromMsg(msg)
					export.SendCwppScanStartToTeams(request, start)
					break
				default:
					logger.Printf("unknown event: %s", msg.Event)
				}
				break
			case "wiki":
				request := &thirdParty.WikiRequest{
					InstanceUrl: msg.InstanceUrl,
					UserName:    msg.UserName,
					Token:       msg.Token,
					SpaceKey:    msg.SpaceKey,
				}
				switch msg.Event {
				case "afterCwppScan":
					result := getScanResultArgsFromMsg(msg)
					export.SendCwppScanResultToWiki(request, result)
					break
				case "beforeCwppScan":
					start := getScanStartArgsFromMsg(msg)
					export.SendCwppScanStartToWiki(request, start)
					break
				default:
					logger.Printf("unknown event: %s", msg.Event)
				}
				break
			case "github":
				request := &thirdParty.GithubRequest{
					Token:      msg.Token,
					Repository: msg.Repository,
				}
				switch msg.Event {
				case "afterCwppScan":
					result := getScanResultArgsFromMsg(msg)
					export.SendCwppScanResultToGithub(request, result)
					break
				case "beforeCwppScan":
					start := getScanStartArgsFromMsg(msg)
					export.SendCwppScanStartToGithub(request, start)
					break
				default:
					logger.Printf("unknown event: %s", msg.Event)
				}
				break
			case "gitlab":
				request := &thirdParty.GitlabRequest{
					Token:     msg.Token,
					ProjectId: msg.ProjectId,
				}
				switch msg.Event {
				case "afterCwppScan":
					result := getScanResultArgsFromMsg(msg)
					export.SendCwppScanResultToGitlab(request, result)
					break
				case "beforeCwppScan":
					start := getScanStartArgsFromMsg(msg)
					export.SendCwppScanStartToGitlab(request, start)
					break
				default:
					logger.Printf("unknown event: %s", msg.Event)
				}
				break
			default:
				logger.Printf("unknown application: %s", msg.Application)
			}
		}
	}()

	logger.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func getScanStartArgsFromMsg(msg Message) *scan.StartInfo {
	args := msg.Args.(map[string]interface{})
	provider := args["provider"].(string)
	userId := args["userId"].(string)
	scanGroupName := args["scanGroupName"].(string)
	keyName := args["keyName"].(string)
	eventTime := args["eventTime"].(string)

	start := &scan.StartInfo{
		Provider:      provider,
		UserId:        userId,
		ScanGroupName: scanGroupName,
		KeyName:       keyName,
		EventTime:     eventTime,
	}
	return start
}

func getScanResultArgsFromMsg(msg Message) *scan.ResultInfo {
	args := msg.Args.(map[string]interface{})
	provider := args["provider"].(string)
	userId := args["userId"].(string)
	scanGroupName := args["scanGroupName"].(string)
	keyName := args["keyName"].(string)
	eventTime := args["eventTime"].(string)
	summary := args["summary"].(map[string]interface{})

	total := summary["total"].(map[string]interface{})
	critical := summary["critical"].(map[string]interface{})
	high := summary["high"].(map[string]interface{})
	medium := summary["medium"].(map[string]interface{})
	low := summary["low"].(map[string]interface{})
	result := &scan.ResultInfo{
		Provider:      provider,
		UserId:        userId,
		ScanGroupName: scanGroupName,
		KeyName:       keyName,
		EventTime:     eventTime,
		ResultSummary: scan.ResultSummary{
			Total: scan.ResultSummaryData{
				Count:      int(total["count"].(float64)),
				Percentage: total["percentage"].(string),
			},
			Critical: scan.ResultSummaryData{
				Count:      int(critical["count"].(float64)),
				Percentage: critical["percentage"].(string),
			},
			High: scan.ResultSummaryData{
				Count:      int(high["count"].(float64)),
				Percentage: high["percentage"].(string),
			},
			Medium: scan.ResultSummaryData{
				Count:      int(medium["count"].(float64)),
				Percentage: medium["percentage"].(string),
			},
			Low: scan.ResultSummaryData{
				Count:      int(low["count"].(float64)),
				Percentage: low["percentage"].(string),
			},
		},
	}
	return result
}
