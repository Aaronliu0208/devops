package main

import (
	"fmt"
	"os"

	"casicloud.com/ylops/marco/config"
	"casicloud.com/ylops/marco/resources"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "uninstall resty package",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("uninstall resty package on %s\n", config.C.Workspace)
		installer := &resources.RestyInstaller{
			BuildDir:     config.C.GetBuildDir(),
			Prefix:       config.C.GetPrefix(),
			BuildOptions: []string{"--with-http_mp4_module"},
		}
		err := installer.Uninstall()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println("uninstall complete")
	},
}

func init() {
	pmCmd.AddCommand(uninstallCmd)
}
