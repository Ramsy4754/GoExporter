# Build Project

Windows

```
batch.bat
```

Linux

```
export GOARCH=arm64
export GOOS=linux
export CC=aarch64-linux-gnu-gcc
export CGO_ENABLED=1
go build -o go_exporter
```

# Rabbit MQ Install

```
sudo apt-get install rabbitmq-server
sudo systemctl enable rabbitmq-server
sudo systemctl start rabbitmq-server.service
```

# Message Format

```
{
    "application": "slack",
    "webhookUrl": "<SLACK_WEBHOOK_URL>",
    "event": "afterCwppScan",
    "args": {
        "provider": "<provider>",
        "userId": "<userId>",
        "scanGroupName": "<scanGroupName>",
        "keyName": "<keyName>", 
        "summary": {
            "total": {
                "count": 72,
                "percentage": "100.00%"
            },
            "critical": {
                "count": 4,
                "percentage": "5.56%"
            },
            "high": {
                "count": 12,
                "percentage": "16.67%"
            },
            "medium": {
                "count": 41,
                "percentage": "56.94"
            },
            "low": {
                "count": 10,
                "percentage": "13.89%",
            }
        }
    }
}

{
    "application": "slack",
    "webhookUrl": "<SLACK_WEBHOOK_URL>",
    "event": "beforeCwppScan",
    "args": {
        "provider": "<provider>",
        "userId": "<userId>",
        "scanGroupName": "<scanGroupName>",
        "keyName": "<keyName>",
        "eventTime": "<eventTime>"
    }
}
```

```
{
    "application": "jira",
    "instanceUrl": "<INSTANCE_URL>",
    "apiKey": "<API_KEY>",
    "projectKey": "<PROJECT_KEY>",
    "event": "beforeCwppScan",
    "args": {
        "provider": "<provider>",
        "userId": "<userId>",
        "scanGroupName": "<scanGroupName>",
        "keyName": "<keyName>",
        "eventTime": "<eventTime>"
    }
}
```

```
{
    "application": "github",
    "event": "beforeCwppScan",
    "token": "<GITHUB_PERSONAL_ACCESS_TOKEN>",
    "repository": "<ACCOUNT/REPO>",
    "args": {
        "provider": "aws",
        "userId": "ramsy4754",
        "scanGroupName": "cwpp scan group 01",
        "keyName": "cwpp key 01",
        "eventTime": "2024-09-09 15:23:13",
    }
}
```

```
{
    "application": "gitlab",
    "event": "beforeCwppScan",
    "projectId": "<GITLAB_PROJECT_ID>",
    "token": "<GITLAB_PERSONAL_ACCESS_TOKEN>",
    "repository": "<GROUP/REPO>",
    "args": {
        "provider": "aws",
        "userId": "ramsy4754",
        "scanGroupName": "cwpp scan group 01",
        "keyName": "cwpp key 01",
        "eventTime": "2024-09-09 15:23:13",
    }
}
```
