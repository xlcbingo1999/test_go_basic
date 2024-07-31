package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xlcbingo1999/test_go_basic/pkg/selectX"
)

var selectDemoCmd = &cobra.Command{
	Use:   "select_demo",
	Short: "Run select_demo",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln("Recover err", err)
			}
		}()

		selectX.RunSelect()
	},
}

func init() {
	rootCmd.AddCommand(selectDemoCmd)
}
