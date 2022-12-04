package cmd

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/webhook"
	"github.com/disgoorg/log"
	"github.com/disgoorg/snowflake/v2"
)

func DiscordSend(Icon string, Col int) error {
	log.SetLevel(log.LevelDebug)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	Cfg, err := loadCfg()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var (
		webhookID    = snowflake.ID(Cfg.DiscordCfg.DiscordWebhookId)
		webhookToken = Cfg.DiscordCfg.DiscordWebhookToken
	)

	fmt.Println(webhookID)
	fmt.Println(webhookToken)

	// construct new webhook client
	client := webhook.New(webhookID, webhookToken)
	defer client.Close(context.TODO())

	// new sync.WaitGroup to await all messages to be sent before shutting down
	var wg sync.WaitGroup
	wg.Add(1)
	go send(&wg, client, Icon, Col)

	// wait for all messages to be sent
	wg.Wait()
	return nil
}

// send(s) a message to the webhook
func send(wg *sync.WaitGroup, client webhook.Client, Icon string, Col int) {
	defer wg.Done()
	Env, err := loadEnv()
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
	)

	StartAtFromUnix := time.Unix(StartAt/1000000, StartAt%1000000).String()
	EndAtFromUnix := time.Unix(EndAt/1000000, EndAt%1000000).String()
	DurarionFromUnix := time.Unix(0, Durarion).String()

	if _, err := client.CreateMessage(discord.NewWebhookMessageCreateBuilder().
		SetEmbeds(
			discord.Embed{
				Title: Icon + Name,
				Color: Col,
				Fields: []discord.EmbedField{
					{
						Name:  "Channel",
						Value: CHName + "/" + CHType,
					},
					{
						Name:  "Time",
						Value: StartAtFromUnix + " ~ " + EndAtFromUnix + " (" + DurarionFromUnix + " )",
					},
					{
						Name:  "Description",
						Value: Description,
					},
					{
						Name:  "RecPath",
						Value: RecPath,
					},
				},
			}).Build(),
	); err != nil {
		log.Fatalf("error: %v", err)
	}
}
