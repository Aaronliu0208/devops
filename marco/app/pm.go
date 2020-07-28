package app

import (
	"casicloud.com/ylops/marco/config"
	"casicloud.com/ylops/marco/pkg/erron"
	"casicloud.com/ylops/marco/pkg/logger"
	"casicloud.com/ylops/marco/pkg/models"
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

// Start cluster
func (pm *PackageManager) Start(cluster *models.Cluster) error {
	binPath, exists := pm.Installer.CheckExists()
	if !exists {
		return erron.New(erron.ErrFileNotFound, "resty not install")
	}
	// build nginx config
	log.Debugln("begin build nginx config")

	log.Debugln("begin start resty")
	ctl := &NginxController{
		BinPath: binPath,
	}

	return ctl.Start()
}
