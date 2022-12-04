package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/slack-go/slack"
)

func Slack(Icon string, Col string) error {
	Env, err := loadEnv()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	Cfg, err := loadCfg()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var (
		Name        = Env.Name
		CHName      = Env.ChannelName
		CHType      = Env.ChannelType
		StartAt     = Env.StartAt
		EndAt       = Env.EndAt
		Durarion    = Env.Durarion
		Description = Env.Description
		RecPath     = Env.RecPath
		SlackKey    = Cfg.SlackCfg.SlackToken
	)

	StartAtFromUnix := time.Unix(StartAt/1000000000, StartAt%1000000000).String()
	EndAtFromUnix := time.Unix(EndAt/1000000000, EndAt%1000000000).String()
	DurarionFromUnix := time.Unix(0, Durarion).String()

	api := slack.New(
		SlackKey,
		slack.OptionDebug(true),
	)
	attachment := slack.Attachment{
		Fallback: Icon + Name,
		Color:    Col,
		Title:    Icon + Name,
		Fields: []slack.AttachmentField{
			{
				Title: "Channel",
				Value: CHName + "/" + CHType,
				Short: false,
			},
			{
				Title: "Time",
				Value: StartAtFromUnix + " ~ " + EndAtFromUnix + " (" + DurarionFromUnix + " )",
				Short: false,
			},
			{
				Title: "Description",
				Value: Description,
				Short: false,
			},
			{
				Title: "RecPath",
				Value: RecPath,
				Short: false,
			},
		},
	}
	channelID, timestamp, err := api.PostMessage(
		Cfg.SlackCfg.Channel,
		slack.MsgOptionAsUser(false),
		slack.MsgOptionAttachments(attachment),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	return nil
}
