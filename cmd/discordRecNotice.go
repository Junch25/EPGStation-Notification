package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// discordRecStartCmd represents the discordRecStart command
var discordRecStartCmd = &cobra.Command{
	Use:   "discordRecStart",
	Short: "Recording start notification command",
	Long:  `This command notifies the start of recording.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := DiscordSend(" :arrow_forward: ", 15158332)
		if err != nil {
			return
		}
		fmt.Println("discordRecStart called")
	},
}

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
	rootCmd.AddCommand(discordRecStartCmd)
	rootCmd.AddCommand(discordRecEndCmd)
	rootCmd.AddCommand(discordRecErrorCmd)
}
