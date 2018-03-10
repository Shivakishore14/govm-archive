package cmd

import (
	"fmt"
	"github.com/shivakishore14/govm/domain"
	"github.com/shivakishore14/govm/engine"
	"github.com/spf13/cobra"
	"os"
	"strings"
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

		versionName := args[0]
		if !strings.HasPrefix(versionName, "go") {
			versionName = "go" + versionName
		}
		for _, x := range remoteVersions {
			if versionName == x.Name {
				version = x
			}
		}
		e := engine.Download(version)
		fmt.Println(e)
	},
}
