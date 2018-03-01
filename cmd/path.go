package cmd

import (
	"fmt"
	"github.com/shivakishore14/govm/engine"
	"github.com/spf13/cobra"
)

var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "finds the path of the go version",
	Long:  `path command finds the path of the go verion given as argument`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("loading .govmrc")
		} else {
			path := engine.Path(args[0])
			fmt.Println(path)
		}
	},
}
