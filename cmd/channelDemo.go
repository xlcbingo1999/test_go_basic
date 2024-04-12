package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xlcbingo1999/test_go_basic/channel"
)

var channelDemoCmd = &cobra.Command{
	Use:   "channel_demo",
	Short: "Run channel_demo",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln("Recover err", err)
			}
		}()

		channel.RunChannelN()
	},
}

func init() {
	rootCmd.AddCommand(channelDemoCmd)
}
