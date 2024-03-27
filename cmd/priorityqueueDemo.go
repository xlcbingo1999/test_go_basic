package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xlcbingo1999/test_go_basic/priorityqueue"
)

var priorityqueueDemoCmd = &cobra.Command{
	Use:   "priorityqueue_demo",
	Short: "Run priorityqueue_demo",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln("Recover err", err)
			}
		}()

		priorityqueue.RunPriorityQueue()
	},
}

func init() {
	rootCmd.AddCommand(priorityqueueDemoCmd)
}
