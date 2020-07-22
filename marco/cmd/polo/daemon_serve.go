package main

import (
	"context"
	"fmt"

	"casicloud.com/ylops/marco/config"
	"casicloud.com/ylops/marco/daemon"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start daemon process to serve cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("serve %s:%d for marco\n", config.C.HTTP.Host, config.C.HTTP.Port)
		ctx := context.Background()
		daemon.Run(ctx)
	},
}

func init() {
	daemonCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntP("http.port", "p", 8080, "port to serve")
	serveCmd.Flags().StringP("http.host", "s", "0.0.0.0", "host to serve")
	viper.BindPFlags(serveCmd.Flags())
}
