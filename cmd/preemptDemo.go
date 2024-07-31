package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xlcbingo1999/test_go_basic/pkg/preemptX"
)

var preemptDemoCmd = &cobra.Command{
	Use:   "preempt_demo",
	Short: "Run preempt_demo",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln("Recover err", err)
			}
		}()

		preemptX.RunPreempt()
	},
}

func init() {
	rootCmd.AddCommand(preemptDemoCmd)
}
