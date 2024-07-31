package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xlcbingo1999/test_go_basic/pkg/semaphoreX"
)

var semaphoreDemoCmd = &cobra.Command{
	Use:   "semaphore_demo",
	Short: "Run semaphore_demo",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln("Recover err", err)
			}
		}()

		semaphoreX.RunSemaphore()
	},
}

func init() {
	rootCmd.AddCommand(semaphoreDemoCmd)
}
