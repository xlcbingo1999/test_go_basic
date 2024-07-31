package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xlcbingo1999/test_go_basic/pkg/futuresX"
)

var futureDemoCmd = &cobra.Command{
	Use:   "future_demo",
	Short: "Run future_demo",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln("Recover err", err)
			}
		}()

		futuresX.RunFuture()
	},
}

func init() {
	rootCmd.AddCommand(futureDemoCmd)
}
