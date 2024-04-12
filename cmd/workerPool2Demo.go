package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xlcbingo1999/test_go_basic/workerPool2"
)

var workerPool2DemoCmd = &cobra.Command{
	Use:   "workerPool2_demo",
	Short: "Run workerPool2_demo",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln("Recover err", err)
			}
		}()

		workerPool2.RunWorkerPool()
	},
}

func init() {
	rootCmd.AddCommand(workerPool2DemoCmd)
}
