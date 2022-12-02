package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// discordRecEndCmd represents the discordRecEnd command
var discordRecEndCmd = &cobra.Command{
	Use:   "discordRecEnd",
	Short: "Recording end notification command",
	Long:  `This command notifies the end of recording.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := DiscordSend(" :white_check_mark: ", 3066993)
		if err != nil {
			return
		}
		fmt.Println("discordRecEnd called")
	},
}

func init() {
	rootCmd.AddCommand(discordRecEndCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// discordRecEndCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// discordRecEndCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
