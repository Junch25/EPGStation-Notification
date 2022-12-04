package cmd

import (
	"os"
	"path/filepath"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type cmdEnv struct {
	ChannelType string `envconfig:"CHANNELTYPE"`
	ChannelID   string `envconfig:"CHANNELID"`
	ChannelName string `envconfig:"CHANNELNAME"`
	StartAt     int64  `envconfig:"STARTAT"`
	EndAt       int64  `envconfig:"ENDAT"`
	Durarion    int64  `envconfig:"DURATION"`
	Name        string `envconfig:"NAME"`
	Description string `envconfig:"DESCRIPTION"`
	Extended    string `envconfig:"EXTENDED"`
	RecPath     string `envconfig:"RECPATH"`
	LogPath     string `envconfig:"LOGPATH"`
	DropCnt     string `envconfig:"DROP_CNT"`
	ErrorCnt    string `envconfig:"ERROR_CNT"`
}

type cmdCfg struct {
	SlackCfg struct {
		SlackToken string `yaml:"slack-token"`
		Channel    string `yaml:"channel"`
	} `yaml:"slack-config"`
	DiscordCfg struct {
		DiscordWebhookToken string `yaml:"discord-webhook-token"`
		DiscordWebhookId    int    `yaml:"discord-webhook"`
	} `yaml:"discord-config"`
}

func loadEnv() (env cmdEnv, err error) {
	if err := envconfig.Process("", &env); err != nil {
		return env, err
	}
	return env, nil
}

func loadCfg() (config cmdCfg, err error) {
	cfg, err := loadYml()
	if err != nil {
		return config, err
	}
	data, err := os.ReadFile(cfg)
	if err != nil {
		return config, err
	}
	err = yaml.UnmarshalStrict([]byte(data), &config)
	if err != nil {
		return config, err
	}
	return config, err
}

func loadYml() (string, error) {
	ymlFilePath, err := os.Executable()
	if err != nil {
		return ymlFilePath, err
	}
	return filepath.Join(filepath.Dir(ymlFilePath), "config.yml"), nil
}
