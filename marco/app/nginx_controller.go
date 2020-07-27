package app

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// NginxController 控制nginx
// 具体功能包括:
// 1. 起停Nginx
// 2. 管理日志
// 3. 管理目录版本
type NginxController struct {
	BinPath    string //path for nginx bin
	Prefix     string // path for nginx prefix
	PidFile    string // path for pid file
	ConfigFile string
}

func (n *NginxController) setDefault() {
	if len(n.PidFile) == 0 {
		n.PidFile = "/var/run/nginx.pid"
	} else {
		abs, err := filepath.Abs(n.PidFile)
		if err == nil {
			n.PidFile = abs
		}
	}

	if len(n.BinPath) == 0 {
		n.BinPath = "/usr/sbin/nginx"
	} else {
		abs, err := filepath.Abs(n.BinPath)
		if err == nil {
			n.BinPath = abs
		}
	}

	if len(n.Prefix) == 0 {
		n.Prefix = "/etc/nginx"
	} else {
		abs, err := filepath.Abs(n.Prefix)
		if err == nil {
			n.Prefix = abs
		}
	}

	if len(n.ConfigFile) == 0 {
		n.ConfigFile = "/etc/nginx/nginx.conf"
	} else {
		abs, err := filepath.Abs(n.ConfigFile)
		if err == nil {
			n.ConfigFile = abs
		}
	}
}

// Start nginx instance
func (n *NginxController) Start() error {
	n.setDefault()
	if _, err := os.Stat(n.PidFile); os.IsNotExist(err) {
		cmd := exec.Command(n.BinPath, "-p", n.Prefix, "-c", n.ConfigFile)
		var outbuf, errbuf bytes.Buffer
		cmd.Stdout = &outbuf
		cmd.Stderr = &errbuf
		err := cmd.Start()
		if err == nil {
			if err := cmd.Wait(); err != nil {
				return err
			}
			if errbuf.Len() != 0 {
				return errors.New(errbuf.String())
			}
		} else {
			return err
		}
		return nil
	}
	return errors.New("process exists")
}

//Stop nginx
func (n *NginxController) Stop() error {
	n.setDefault()
	if _, err := os.Stat(n.PidFile); os.IsNotExist(err) {
		return err
	}

	cmd := exec.Command(n.BinPath, "-p", n.Prefix, "-c", n.ConfigFile, "-s", "stop")
	var outbuf, errbuf bytes.Buffer
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf
	err := cmd.Start()
	if err == nil {
		if err := cmd.Wait(); err != nil {
			return err
		}
		if _, err := os.Stat(n.PidFile); os.IsNotExist(err) {
			return nil
		}

		return errors.New("stop failed")
	}
	return err
}

//Reload nginx
func (n *NginxController) Reload() error {
	n.setDefault()
	cmd := exec.Command(n.BinPath, "-s", "reload", "-c", n.ConfigFile, "-p", n.Prefix)
	var outbuf, errbuf bytes.Buffer
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf
	err := cmd.Start()
	if err == nil {
		if err := cmd.Wait(); err != nil {
			return err
		}

		if errbuf.Len() > 0 {
			return errors.New(errbuf.String())
		}
	}

	return nil
}

//Test nginx
func (n *NginxController) Test() (bool, error) {
	n.setDefault()
	cmd := exec.Command(n.BinPath, "-c", n.ConfigFile, "-p", n.Prefix, "-t")
	var outbuf, errbuf bytes.Buffer
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf
	err := cmd.Start()
	if err != nil {
		return false, err
	}

	if err := cmd.Wait(); err != nil {
		return false, err
	}

	if errbuf.Len() != 0 {
		return strings.Contains(errbuf.String(), "syntax is ok"), nil
	}
	return strings.Contains(outbuf.String(), "syntax is ok"), nil
}
