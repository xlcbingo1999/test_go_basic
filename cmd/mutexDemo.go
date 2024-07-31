package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xlcbingo1999/test_go_basic/pkg/mutexDemo"
)

var mutexDemoCmd = &cobra.Command{
	Use:   "mutex_demo",
	Short: "Run mutex_demo",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln("Recover err", err)
			}
		}()

		mutexDemo.RunMutexDemo()
	},
}

func init() {
	mutexDemoCmd.Flags().StringVarP(&mutexDemo.ServePort, "mutex-serve-port", "", "6060", "mutex demo serve port")
	rootCmd.AddCommand(mutexDemoCmd)
}
