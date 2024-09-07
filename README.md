# Build Project
```
batch.bat
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