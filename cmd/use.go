package cmd

import (
	"fmt"
	"github.com/shivakishore14/govm/engine"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Lets you switch to another go version",
	Long:  `Use command lets you switch between multiple go versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("loading .govmrc")
		} else {
			engine.Use(args[0])
		}
	},
}
