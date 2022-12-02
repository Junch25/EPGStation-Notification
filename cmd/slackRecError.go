package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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
	rootCmd.AddCommand(slackRecErrorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// slackRecErrorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// slackRecErrorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
