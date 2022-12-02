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

func init() {
	rootCmd.AddCommand(discordRecStartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// discordRecStartCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// discordRecStartCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
