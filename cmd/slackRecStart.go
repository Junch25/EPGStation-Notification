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

func init() {
	rootCmd.AddCommand(slackRecStartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// slackRecStartCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// slackRecStartCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
