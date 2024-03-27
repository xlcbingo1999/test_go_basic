package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xlcbingo1999/test_go_basic/httpgroup5"
)

var httpgroup5DemoCmd = &cobra.Command{
	Use:   "httpgroup5_demo",
	Short: "Run httpgroup5_demo",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln("Recover err", err)
			}
		}()

		httpgroup5.RunHttpGroup5()
	},
}

func init() {
	rootCmd.AddCommand(httpgroup5DemoCmd)
}
