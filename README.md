# EPGStation-Notification
EPGStationの録画開始・終了・エラーのSlack・Discord通知です。

## CI/CD
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/Junch25/EPGStation-Notification/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/Junch25/EPGStation-Notification/tree/main)

## 導入手順
### Slackアプリを作成
URL: https://api.slack.com

参考: https://api.slack.com/lang/ja-jp

### スクリプト配置
```shell script
$ git clone https://github.com/Junch25/EPGStation-Notification.git
```

### Slackの設定
```shell script
# 編集
$ cd epgstation_notification
$ vim bin/config.yml
slack-config:
  slack-token: "SLACK_API_TOKEN"
  channel: "SLACK_CHANNEL_ID"

discord-config:
  discord-webhook-token: "DISCORD_API_TOKEN"
  discord-webhook: "DISCORD_API_WEBHOOK"
```

### EPGStationへ設定追加
```shell script
$ vim /path/to/config/config.yml
# Slack
---
recordingStartCommand: "/path/to/bin/epgstation-notification slackRecStart"
recordingFinishCommand: "/path/to/bin/epgstation-notification slackRecEnd"
recordingFailedCommand: "/path/to/bin/epgstation-notification slackRecError"

# Discord
---
recordingStartCommand: "/path/to/bin/epgstation-notification discordRecStart"
recordingFinishCommand: "/path/to/bin/epgstation-notification discordRecEnd"
recordingFailedCommand: "/path/to/bin/epgstation-notification discordRecError"
```
### EPGStation再起動
```shell script
$ sudo pm2 restart epgstation
```

### Build

`EPGStation-Notification`というバイナリファイルできればOK
```shell script
$ cd epgstation_notification
$ GOOS=linux GOARCH=amd64 go build -o "bin/epgstation-notification" main.go
$ ls bin
  epgstation-notification
```

## License
[Apache License 2.0](https://github.com/Junch25/EPGStation-Notification/blob/main/LICENSE)
