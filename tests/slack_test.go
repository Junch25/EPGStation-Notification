package cmd

import (
	"epgstation_notification/cmd"
	"testing"
)

func TestSlack(t *testing.T) {
	type args struct {
		Icon string
		Col  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := cmd.Slack(tt.args.Icon, tt.args.Col); (err != nil) != tt.wantErr {
				t.Errorf("Slack() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
