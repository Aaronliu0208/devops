package config

import (
	"fmt"
	"os"
	"sync"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	// C static config
	C    = new(Config)
	once sync.Once
)

// Config 程序配置类
type Config struct {
	RunMode string
	HTTP    HTTP `yaml:"http,omitempty"`
	Log     Log  `yaml:"log,omitempty"`
}

// HTTP http config
type HTTP struct {
	Host            string `yaml:"host,omitempty"`
	Port            int    `yaml:"port,omitempty"`
	CertFile        string `yaml:"cert_file,omitempty"`
	KeyFile         string `yaml:"key_file,omitempty"`
	ShutdownTimeout int    `yaml:"shutdown_timeout,omitempty"`
}

// LogHook 日志钩子
type LogHook string

// Log 日志配置参数
type Log struct {
	Level         int
	Format        string
	Output        string
	OutputFile    string
	EnableHook    bool
	HookLevels    []string
	Hook          LogHook
	HookMaxThread int
	HookMaxBuffer int
}

// Load the config file from the current directory and marshal
// into the conf config struct.
func Load(cfgFile string) {
	once.Do(func() {
		if cfgFile != "" {
			// Use config file from the flag.
			viper.SetConfigFile(cfgFile)
		} else {
			// Find home directory.
			home, err := homedir.Dir()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Search config in home directory with name ".ylops" (without extension).
			viper.AddConfigPath(".")
			viper.AddConfigPath(home)
			viper.SetConfigName("config")
		}

		viper.AutomaticEnv() // read in environment variables that match

		// If a config file is found, read it in.
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Printf("%v", err)
		}

		err = viper.Unmarshal(C)
		if err != nil {
			fmt.Printf("unable to decode into config struct, %v", err)
		}
	})
}
