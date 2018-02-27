package cmd

import (
	"github.com/shivakishore14/govm/engine"
	"github.com/spf13/cobra"
	"log"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall a golang version",
	Long:  `Uninstalls Go given the version given`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		if len(args) == 0 {
			log.Println("please pecify a version to uninstall")
			return
		}
		if err := engine.Uninstall(args[0]); err != nil {
			log.Println(err)
			return
		}

		log.Println("Successfuly removed ", args[0])

	},
}
