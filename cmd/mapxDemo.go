package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xlcbingo1999/test_go_basic/pkg/mapx"
)

var mapxDemoCmd = &cobra.Command{
	Use:   "mapx_demo",
	Short: "Run mapx_demo",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln("Recover err", err)
			}
		}()

		mapx.RunRWMap()
	},
}

func init() {
	rootCmd.AddCommand(mapxDemoCmd)
}
