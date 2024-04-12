package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xlcbingo1999/test_go_basic/epoll"
)

var epollDemoCmd = &cobra.Command{
	Use:   "epoll_demo",
	Short: "Run epoll_demo",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln("Recover err", err)
			}
		}()

		epoll.RunTCP()
	},
}

func init() {
	rootCmd.AddCommand(epollDemoCmd)
}
