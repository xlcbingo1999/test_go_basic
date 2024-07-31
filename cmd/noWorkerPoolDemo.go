package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xlcbingo1999/test_go_basic/pkg/noWorkerPool"
)

var noWorkerPoolDemoCmd = &cobra.Command{
	Use:   "noWorkerPool_demo",
	Short: "Run noWorkerPool_demo",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln("Recover err", err)
			}
		}()

		noWorkerPool.RunNoWorkerPool()
	},
}

func init() {
	rootCmd.AddCommand(noWorkerPoolDemoCmd)
}
