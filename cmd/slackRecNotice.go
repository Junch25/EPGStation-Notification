package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// slackRecStartCmd represents the slackRecStart command
var slackRecStartCmd = &cobra.Command{
	Use:   "slackRecStart",
	Short: "Recording start notification command",
	Long:  `This command notifies the start of recording.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := Slack(" :arrow_forward: ", "danger")
		if err != nil {
			return
		}
		fmt.Println("slackRecStart called")
	},
}

// slackRecEndCmd represents the slackRecEnd command
var slackRecEndCmd = &cobra.Command{
	Use:   "slackRecEnd",
	Short: "Recording end notification command",
	Long:  `This command notifies the end of recording.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := Slack(" :white_check_mark: ", "good")
		if err != nil {
			return
		}
		fmt.Println("slackRecEnd called")
	},
}

// slackRecErrorCmd represents the slackRecError command
var slackRecErrorCmd = &cobra.Command{
	Use:   "slackRecError",
	Short: "Recording error notification command",
	Long:  `This command notifies you of recording errors.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := Slack(" :x: ", "warning")
		if err != nil {
			return
		}
		fmt.Println("slackRecError called")
	},
}

func init() {
	rootCmd.AddCommand(slackRecStartCmd)
	rootCmd.AddCommand(slackRecEndCmd)
	rootCmd.AddCommand(slackRecErrorCmd)
}
