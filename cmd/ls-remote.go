package cmd

import (
	"github.com/shivakishore14/govm/engine"
	"github.com/spf13/cobra"
	"fmt"
)

var lsRemoteCmd = &cobra.Command{
	Use:   "ls-remote",
	Short: "Display all the versions available",
	Long:  `Display all the versions of Go available for download`,
	Run: func(cmd *cobra.Command, args []string) {
		remoteVersions := engine.RemoteList()
		//fmt.Println(remoteVersions)
		for _, x := range remoteVersions {
			fmt.Println(x.Name, x.DownloadLink)
		}
	},
}
