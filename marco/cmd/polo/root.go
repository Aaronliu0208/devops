package main

import (
	"fmt"
	"os"

	"casicloud.com/ylops/marco/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "polo",
	Short: "polo a command line to manage marco instance",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.marco.yaml)")

	rootCmd.PersistentFlags().StringP("workspace", "w", "~/marco", "workspace for install pm package")
	viper.BindPFlag("workspace", rootCmd.PersistentFlags().Lookup("workspace"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	config.InitConfig(cfgFile)
}
