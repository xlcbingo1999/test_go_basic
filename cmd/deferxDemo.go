package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xlcbingo1999/test_go_basic/pkg/deferx"
)

var deferxDemoCmd = &cobra.Command{
	Use:   "deferx_demo",
	Short: "Run deferx_demo",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln("Recover err", err)
			}
		}()

		deferx.RunDefer()
	},
}

func init() {
	rootCmd.AddCommand(deferxDemoCmd)
}
