package cmd

import (
	"fmt"
	"log"

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
		Description = Env.Description
		RecPath     = Env.RecPath
		LogPath     = Env.LogPath
		SlackKey    = Cfg.SlackCfg.SlackToken
	)

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
