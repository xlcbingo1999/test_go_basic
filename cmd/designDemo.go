package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xlcbingo1999/test_go_basic/pkg/design"
)

var designDemoCmd = &cobra.Command{
	Use:   "design_demo",
	Short: "Run design_demo",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln("Recover err", err)
			}
		}()

		design.RunAllDesign()
	},
}

func init() {
	rootCmd.AddCommand(designDemoCmd)
}
