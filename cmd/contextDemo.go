package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xlcbingo1999/test_go_basic/contextX"
)

var contextDemoCmd = &cobra.Command{
	Use:   "context_demo",
	Short: "Run context_demo",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln("Recover err", err)
			}
		}()

		// contextX.RunNormalClose()
		// contextX.RunContextExample1()
		contextX.RunContextExample3()
	},
}

func init() {
	rootCmd.AddCommand(contextDemoCmd)
}