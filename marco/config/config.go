package config

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	// C config instance must init on start
	C    = new(Config)
	once sync.Once
)

// Config of marco
type Config struct {
	Mode      string `mapstructure:"mode"`
	Workspace string `mapstructure:"workspace,omitempty"`
	HTTP      HTTP   `mapstructure:"http,omitempty"`
	Log       Log    `mapstructure:"log,omitempty"`
	Git       Git    `mapstructure:"git,omitempty"`
}

//GetBuildDir get build dir
func (c Config) GetBuildDir() string {
	return filepath.Join(c.Workspace, "build")
}

//GetPrefix get resty prefix
func (c Config) GetPrefix() string {
	return filepath.Join(c.Workspace, "app")
}

func (c Config) GetSnapshotDir() string {
	return filepath.Join(c.Workspace, "snaps")
}

//GetNginxBinPath get nginx bin path
func (c Config) GetNginxBinPath() string {
	return filepath.Join(c.Workspace, "app/nginx/sbin/nginx")
}

func (c Config) GetNginxConfigPath() string {
	return filepath.Join(c.Workspace, "app/nginx/conf/nginx.conf")
}

//GetLogDir get log prefix dir
func (c Config) GetLogDir() string {
	return filepath.Join(c.Workspace, "logs")
}

//GetTempDir get temp file folder
func (c Config) GetTempDir() string {
	return filepath.Join(c.Workspace, ".tmp")
}

func (c Config) GetPid() string {
	return filepath.Join(c.GetLogDir(), "nginx.pid")
}

func (c Config) EnsureDirectoryExists() error {
	if err := os.MkdirAll(c.GetBuildDir(), 0777); err != nil {
		return err
	}

	if err := os.MkdirAll(c.GetLogDir(), 0777); err != nil {
		return err
	}

	if err := os.MkdirAll(c.GetPrefix(), 0777); err != nil {
		return err
	}

	if err := os.MkdirAll(c.GetTempDir(), 0777); err != nil {
		return err
	}
	return nil
}

// HTTP config for http server
type HTTP struct {
	Host            string `mapstructure:"host,omitempty"`
	Port            int    `mapstructure:"port,omitempty"`
	EnableSSL       bool   `mapstructure:"enable_ssl,omitempty"`
	CertFile        string `mapstructure:"cert_file,omitempty"`
	KeyFile         string `mapstructure:"key_file,omitempty"`
	ShutdownTimeout int    `mapstructure:"shutdown_timeout,omitempty"`
}

// Log config for log
type Log struct {
	//same as logrus level
	DEBUG bool `mapstructure:"level,omitempty"`
	// output type like stdout/stderr/file
	Output string `mapstructure:"output,omitempty"`
	// File name for app log
	File string `mapstructure:"file,omitempty"`
}

// Git config for git
type Git struct {
	CertFile string `mapstructure:"cert_file,omitempty"`
	KeyFile  string `mapstructure:"key_file,omitempty"`
	RepoURL  string `mapstructure:"repo_url,omitempty"`
}

// Postgres postgres配置参数
type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// DSN 数据库连接串
func (a Postgres) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		a.Host, a.Port, a.User, a.DBName, a.Password, a.SSLMode)
}

//InitConfig 初始化数据库，就运行一次
func InitConfig(cfgFile string) {
	once.Do(func() {
		config, err := LoadConfigFile(cfgFile)
		if err != nil {
			panic(err)
		}
		C = config
	})
}

// LoadConfig by string
func LoadConfig(conf []byte) (*Config, error) {
	config := &Config{}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv() // read in environment variables that match
	viper.ReadConfig(bytes.NewBuffer(conf))
	err := viper.Unmarshal(config)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
		return config, err
	}

	return config, nil
}

//LoadConfigFile load config file, if error panic
func LoadConfigFile(cfgFile string) (*Config, error) {
	config := &Config{}
	viper.SetConfigType("yaml")
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			return config, err
		}

		// Search config in home directory with name "config" (without extension).
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			fmt.Println(err)
			return config, err
		}
		viper.AddConfigPath(dir)
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("%v", err)
		return config, err
	}

	err = viper.Unmarshal(config)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
		return config, err
	}

	return config, nil
}
