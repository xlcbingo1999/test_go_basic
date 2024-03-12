package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xlcbingo1999/test_go_basic/productConsume"
)

var productConsumeDemoCmd = &cobra.Command{
	Use:   "productConsume_demo",
	Short: "Run productConsume_demo",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln("Recover err", err)
			}
		}()

		productConsume.RunProductConsume()
	},
}

func init() {
	rootCmd.AddCommand(productConsumeDemoCmd)
}
