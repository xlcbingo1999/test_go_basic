package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "test_go_basic",
	Short: "run test_go_basic project",
}

func Execute() error {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln("fail to start cobra execute: ", err)
	}
	return err
}
