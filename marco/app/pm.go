package app

import (
	"fmt"
	"os"

	"casicloud.com/ylops/marco/config"
	"casicloud.com/ylops/marco/pkg/erron"
	"casicloud.com/ylops/marco/pkg/logger"
	"casicloud.com/ylops/marco/pkg/models"
	"casicloud.com/ylops/marco/pkg/utils"
	"casicloud.com/ylops/marco/resources"
)

var log = logger.Get("pakcage manager", nil)

//PackageManager 管理nginx 包括安装，以及启停管理
type PackageManager struct {
	Workspace string
	Installer *resources.RestyInstaller
	Config    *config.Config
}

//NewPM new package manager
func NewPM(config *config.Config) *PackageManager {
	pm := &PackageManager{
		Workspace: config.Workspace,
		Installer: &resources.RestyInstaller{
			BuildDir:     config.GetBuildDir(),
			Prefix:       config.GetPrefix(),
			BuildOptions: []string{"--with-http_mp4_module"},
		},
		Config: config,
	}

	return pm
}
func (pm *PackageManager) Init() error {
	// create all work dirs
	if err := pm.Config.EnsureDirectoryExists(); err != nil {
		return err
	}
	// init repo
	return nil
}

// Start cluster
func (pm *PackageManager) Start(cluster *models.Cluster) error {
	binPath, exists := pm.Installer.CheckExists()
	if !exists {
		return erron.New(erron.ErrFileNotFound, "resty not install")
	}
	// build nginx config
	log.Debugln("begin build nginx config")
	config, err := cluster.GenerateConfig()
	if err != nil {
		return err
	}

	// create temp config
	tmpDir := pm.Config.GetTempDir()
	if _, err := os.Stat(tmpDir); os.IsNotExist(err) {
		os.MkdirAll(tmpDir, 0777)
	}
	tmpFIlePath, err := utils.GenerateTempFileName(tmpDir, "nginx.conf.*")
	if err != nil {
		return err
	}
	if err = utils.AppendFileString(tmpFIlePath, config); err != nil {
		return err
	}

	log.Debugln("begin start resty")
	ctl := &NginxController{
		BinPath:    binPath,
		Prefix:     pm.Config.Workspace, //这里要求Prefix在Workspace目录下
		ConfigFile: tmpFIlePath,
		PidFile:    pm.Config.GetPid(),
	}

	ok, err := ctl.Test()
	if err != nil {
		return err
	}

	if ok {
		//mv temp file to config folder
		err := utils.MoveFile(tmpFIlePath, pm.Config.GetNginxConfigPath(), true)
		if err != nil {
			return err
		}
		ctl.ConfigFile = pm.Config.GetNginxConfigPath()
		return ctl.Start()
	}

	return fmt.Errorf("start cluster fail")
}
