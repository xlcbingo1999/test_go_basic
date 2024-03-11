package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xlcbingo1999/test_go_basic/singleton"
)

var ginDemoCmd = &cobra.Command{
	Use:   "singleton_demo",
	Short: "Run singleton_demo",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln("Recover err", err)
			}
		}()

		singleton.RunTest()
	},
}

func init() {
	rootCmd.AddCommand(ginDemoCmd)
}
