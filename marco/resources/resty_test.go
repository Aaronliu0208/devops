package resources

import (
	"os"
	"testing"
)

func TestRestyInstaller(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	installer := &RestyInstaller{
		WorkDir: path + "/test",
	}

	err = installer.Extract()
	if err != nil {
		t.Fatal(err)
	}
	os.RemoveAll(path + "/test")
}
