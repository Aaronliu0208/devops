package resources

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"casicloud.com/ylops/marco/pkg/logger"
)

var log = logger.Get("RestyInstaller", nil)

//RestyInstaller install resty on disk
type RestyInstaller struct {
	BuildDir     string
	Prefix       string
	BuildOptions []string
}

func (r *RestyInstaller) Uninstall() error {
	err := os.RemoveAll(r.BuildDir)
	if err != nil {
		return err
	}
	return os.RemoveAll(r.Prefix)
}

// CheckExists check resty exists
func (r *RestyInstaller) CheckExists() (string, bool) {
	var list []string
	err := filepath.Walk(r.Prefix, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}

		//正则匹配路径名和需要查找的文件名
		ok, _ := regexp.MatchString("nginx$", path)
		if ok && IsExecAll(f.Mode()) {
			list = append(list, path)
		}
		return nil
	})
	exists := err == nil && len(list) != 0
	if exists {
		return list[0], true
	}

	return "", false
}

// Install openresty
func (r *RestyInstaller) Install(ctx context.Context) error {
	// 清理build dir
	os.RemoveAll(r.BuildDir)
	//解压
	err := r.Extract()
	if err != nil {
		return err
	}
	//chmod
	configurePath := filepath.Join(r.BuildDir, "configure")
	if _, err := os.Stat(configurePath); os.IsNotExist(err) {
		return err
	}
	err = os.Chmod(configurePath, 0755)
	if err != nil {
		return err
	}
	//run config
	tmpfile, err := ioutil.TempFile(r.BuildDir, "*.out")
	if err != nil {
		return err
	}
	mwriter := io.MultiWriter(tmpfile, os.Stdout)
	/*
			./configure --prefix=${RESTY_PREFIX} \
		  --with-pcre-jit \
		  --with-cc-opt="-I/usr/local/include" \
		  --with-ld-opt="-L/usr/local/lib" \
		  --with-http_stub_status_module \
		  --with-http_mp4_module \
	*/
	builder := strings.Builder{}
	builder.WriteString("./configure ")
	builder.WriteString(fmt.Sprintf("--prefix=%s ", r.Prefix))
	builder.WriteString("--with-pcre-jit ")
	builder.WriteString("--with-cc-opt=\"-I/usr/local/include\" ")
	builder.WriteString("--with-ld-opt=\"-L/usr/local/lib\" ")
	builder.WriteString("--with-http_stub_status_module ")
	for _, v := range r.BuildOptions {
		builder.WriteString(v)
		builder.WriteString(" ")
	}
	cmd := exec.CommandContext(ctx, "/bin/bash", "-c", builder.String())
	cmd.Stderr = mwriter
	cmd.Stdout = mwriter
	cmd.Dir = r.BuildDir
	err = cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		return err
	}

	// change install file mode
	installPath := r.BuildDir + "/build/install"
	if _, err := os.Stat(installPath); os.IsNotExist(err) {
		return err
	}
	err = os.Chmod(installPath, 0755)
	if err != nil {
		return err
	}
	// make and make install
	cmd = exec.CommandContext(ctx, "/bin/bash", "-c", "make && make install")
	cmd.Stderr = mwriter
	cmd.Stdout = mwriter
	cmd.Dir = r.BuildDir
	err = cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		return err
	}
	return nil
}

// Extract extract resty tar.gz
func (r *RestyInstaller) Extract() error {
	parentDir := "openresty-1.17.8.2"
	assertName := "openresty-1.17.8.2.tar.gz"
	data, err := Asset(assertName)
	if err != nil {
		return err
	}

	gzf, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(gzf)
	for true {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("extract: Next() failed: %s", err.Error())
		}

		name := strings.Replace(header.Name, parentDir, "", -1)
		path := filepath.Join(r.BuildDir, name)
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(path, 0755); err != nil {
				if os.IsExist(err) {
					continue
				}
				return fmt.Errorf("extract: Mkdir() failed: %s", err.Error())
			}
		case tar.TypeReg:
			outFile, err := os.Create(path)
			if err != nil {
				return fmt.Errorf("extract: Create() failed: %s", err.Error())
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return fmt.Errorf("extract: Copy() failed: %s", err.Error())
			}
			outFile.Close()

		default:
			return fmt.Errorf(
				"extract: uknown type: %s in %s",
				string(header.Typeflag),
				name)
		}
	}
	return nil
}
