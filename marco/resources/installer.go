package resources

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//RestyInstaller install resty on disk
type RestyInstaller struct {
	WorkDir      string
	Prefix       string
	BuildOptions []string
}

// Install openresty
func (r *RestyInstaller) Install(ctx context.Context) error {
	//解压
	//chmod
	//run config
	// make and make install
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
		path := filepath.Join(r.WorkDir, name)
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
