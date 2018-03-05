package cmd

import (
	"fmt"
	"github.com/shivakishore14/govm/domain"
	"github.com/shivakishore14/govm/engine"
	"github.com/spf13/cobra"
	"os"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a golang version",
	Long:  `Installs Go given the version to be installed`,
	Run: func(cmd *cobra.Command, args []string) {
		hostOs := os.Getenv("GOVMOS")
		hostArch := os.Getenv("GOVMARCH")
		fmt.Println(hostOs, hostArch)
		if hostOs == "" || hostArch == "" {
			fmt.Println("please check configuration")
		}
		remoteVersions := engine.RemoteList(hostOs, hostArch)
		version := domain.Version{}
		for _, x := range remoteVersions {
			if args[0] == x.Name {
				version = x
			}
		}
		e := engine.Download(version)
		fmt.Println(e)
	},
}
