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

			switch msg.Application {
			case "slack":
				switch msg.Event {
				case "afterCwppScan":
					request := &thirdParty.SlackRequest{
						WebhookUrl: msg.WebhookUrl,
					}

					result := &scan.Result{
						ScanType:        "",
						Vulnerabilities: nil,
					}
					export.SendCwppScanResultToSlack(request, result)
					break
				case "beforeCwppScan":
					request := &thirdParty.SlackRequest{
						WebhookUrl: msg.WebhookUrl,
					}

					args := msg.Args.(map[string]interface{})
					provider := args["provider"].(string)
					userId := args["userId"].(string)
					scanGroupName := args["scanGroupName"].(string)
					keyName := args["keyName"].(string)
					start := &scan.StartInfo{
						Provider:      provider,
						UserId:        userId,
						ScanGroupName: scanGroupName,
						KeyName:       keyName,
					}

					export.SendCwppScanStartToSlack(request, start)
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
