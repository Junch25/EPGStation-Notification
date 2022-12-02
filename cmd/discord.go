package cmd

import (
	"context"
	"fmt"
	"sync"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/webhook"
	"github.com/disgoorg/log"
	"github.com/disgoorg/snowflake/v2"
)

func DiscordSend(Icon string, Col int) error {
	log.SetLevel(log.LevelDebug)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cfg, err := loadCfg()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var (
		webhookID    = snowflake.ID(cfg.DiscordCfg.DiscordWebhookId)
		webhookToken = cfg.DiscordCfg.DiscordWebhookToken
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
	env, err := loadEnv()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var (
		Name        = env.Name
		CHName      = env.ChannelName
		CHType      = env.ChannelType
		Description = env.Description
		RecPath     = env.RecPath
		LogPath     = env.LogPath
	)

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
