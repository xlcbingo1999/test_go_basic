package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xlcbingo1999/test_go_basic/pkg/barrierX"
)

var barrierDemoCmd = &cobra.Command{
	Use:   "barrier_demo",
	Short: "Run barrier_demo",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln("Recover err", err)
			}
		}()

		barrierX.RunBarrier()
	},
}

func init() {
	rootCmd.AddCommand(barrierDemoCmd)
}
