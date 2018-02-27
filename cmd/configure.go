package cmd

import "github.com/spf13/cobra"

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "configure Govm",
	Long:  `configure Govm with initial data and this is to be run as a part of installation`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

