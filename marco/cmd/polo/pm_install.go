package main

import (
	"context"
	"fmt"
	"os"

	"casicloud.com/ylops/marco/config"
	"casicloud.com/ylops/marco/resources"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install resty package",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("install resty package on %s\n", config.C.Workspace)
		installer := &resources.RestyInstaller{
			BuildDir:     config.C.GetBuildDir(),
			Prefix:       config.C.GetPrefix(),
			BuildOptions: []string{"--with-http_mp4_module"},
		}
		ctx := context.Background()
		err := installer.Install(ctx)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	pmCmd.AddCommand(installCmd)
}
