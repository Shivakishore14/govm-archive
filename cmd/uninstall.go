package cmd

import (
	"github.com/shivakishore14/govm/engine"
	"github.com/spf13/cobra"
	"fmt"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall a golang version",
	Long:  `Uninstalls Go given the version given`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("please pecify a version to uninstall")
			return
		}
		if err := engine.Uninstall(args[0]); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Successfuly removed ", args[0])

	},
}
