package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xlcbingo1999/test_go_basic/workerPool"
)

var workerPoolDemoCmd = &cobra.Command{
	Use:   "workerPool_demo",
	Short: "Run workerPool_demo",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln("Recover err", err)
			}
		}()

		workerPool.RunWorkerPool()
	},
}

func init() {
	rootCmd.AddCommand(workerPoolDemoCmd)
}
