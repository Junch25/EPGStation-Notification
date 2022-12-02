package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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

func init() {
	rootCmd.AddCommand(slackRecEndCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// slackRecEndCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// slackRecEndCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
