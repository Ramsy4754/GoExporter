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
			var msgMap map[string]interface{}
			if err := json.Unmarshal(d.Body, &msgMap); err != nil {
				logger.Printf("error parsing JSON: %s", err)
				continue
			}

			app, ok := msgMap["application"]
			if !ok {
				logger.Printf("error parsing application")
			}

			result, ok := msgMap["result"]
			if !ok {
				logger.Printf("error parsing result")
			}

			resultMap, ok := result.(map[string]interface{})
			if !ok {
				logger.Printf("error parsing result")
			}

			scanType, ok := resultMap["scanType"].(string)
			if !ok {
				logger.Printf("error parsing scanType")
			}

			vulnerabilities, ok := resultMap["Vulnerabilities"].([]scan.Vulnerability)
			if !ok {
				logger.Printf("error parsing vulnerabilities")
			}

			scanResult := &scan.Result{
				ScanType:        scanType,
				Vulnerabilities: vulnerabilities,
			}

			switch app {
			case "slack":
				webhookUrl, ok := msgMap["webhookUrl"].(string)
				if !ok {
					logger.Printf("error parsing webhookUrl: %s", err)
					continue
				}
				request := &thirdParty.SlackRequest{
					WebhookUrl: webhookUrl,
				}
				export.SendMessageToSlack(request, scanResult)
				break
			}
		}
	}()

	logger.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
