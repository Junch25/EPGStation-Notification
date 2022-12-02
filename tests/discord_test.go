package cmd

import (
	"epgstation_notification/cmd"
	"testing"
)

func Test_send(t *testing.T) {
	type args struct {
		Icon string
		Col  int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd.DiscordSend(tt.args.Icon, tt.args.Col)
		})
	}
}
