package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// discordRecErrorCmd represents the discordRecError command
var discordRecErrorCmd = &cobra.Command{
	Use:   "discordRecError",
	Short: "Recording error notification command",
	Long:  `This command notifies you of recording errors.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := DiscordSend(" :x: ", 16776960)
		if err != nil {
			return
		}
		fmt.Println("discordRecError called")
	},
}

func init() {
	rootCmd.AddCommand(discordRecErrorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// discordRecErrorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// discordRecErrorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
