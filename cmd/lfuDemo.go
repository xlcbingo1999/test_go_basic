package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xlcbingo1999/test_go_basic/pkg/lfu"
)

var lfuDemoCmd = &cobra.Command{
	Use:   "lfu_demo",
	Short: "Run lfu_demo",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln("Recover err", err)
			}
		}()

		lfu.RunLFU()
	},
}

func init() {
	rootCmd.AddCommand(lfuDemoCmd)
}
