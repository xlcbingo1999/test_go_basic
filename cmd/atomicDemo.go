package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xlcbingo1999/test_go_basic/pkg/atomicX"
)

var atomicDemoCmd = &cobra.Command{
	Use:   "atomic_demo",
	Short: "Run atomic_demo",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln("Recover err", err)
			}
		}()

		atomicX.RunAtomic()
	},
}

func init() {
	rootCmd.AddCommand(atomicDemoCmd)
}
