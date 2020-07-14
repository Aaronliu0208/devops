package resources

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"log"
	"os"
	"testing"
)

func TestRestyAssert(t *testing.T) {
	os.RemoveAll("openresty-1.17.8.2")
	data, err := Asset("openresty-1.17.8.2.tar.gz")
	if err != nil {
		// Asset was not found.
		t.Fatal(err)
	}

	gzf, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	tarReader := tar.NewReader(gzf)
	for true {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("ExtractTarGz: Next() failed: %s", err.Error())
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(header.Name, 0755); err != nil {
				t.Fatalf("ExtractTarGz: Mkdir() failed: %s", err.Error())
			}
		case tar.TypeReg:
			outFile, err := os.Create(header.Name)
			if err != nil {
				t.Fatalf("ExtractTarGz: Create() failed: %s", err.Error())
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				t.Fatalf("ExtractTarGz: Copy() failed: %s", err.Error())
			}
			outFile.Close()

		default:
			t.Fatalf(
				"ExtractTarGz: uknown type: %s in %s",
				string(header.Typeflag),
				header.Name)
		}
	}
	// os.RemoveAll("openresty-1.17.8.2")
}
