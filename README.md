# Message Format

## Slack

```
{
    "applicatoin": "slack",
    "webhookUrl": "<SLACK_WEBHOOK_URL>",
    "scanType": "image_scan",
    "vulnerabilities": [
        {"cve": "CVE-2021-44228", "severity": "Critical", "description": "Log4j vulnerability"},
        {"cve": "CVE-2022-22965", "severity": "High", "description": "Spring4Shell vulnerability"},
    ],
}
```