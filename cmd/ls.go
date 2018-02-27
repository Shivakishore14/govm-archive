package cmd

import (
	"github.com/shivakishore14/govm/engine"
	"github.com/spf13/cobra"
	"log"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Display the versions installed",
	Long:  `Display all the versions of Go installed`,
	Run: func(cmd *cobra.Command, args []string) {
		remoteVersions := engine.LocalList()
		//log.Println(remoteVersions)
		for _, x := range remoteVersions {
			log.Println(x.Name, x.DownloadLink)
		}
	},
}
