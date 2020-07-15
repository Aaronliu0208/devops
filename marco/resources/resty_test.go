package resources

import (
	"io"
	"os"
	"os/exec"
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

func TestMultiStdout(t *testing.T) {
	// Logging capability
	f, err := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()
	mwriter := io.MultiWriter(f, os.Stdout)
	cmd := exec.Command("ls")
	cmd.Stderr = mwriter
	cmd.Stdout = mwriter
	err = cmd.Run() //blocks until sub process is complete
	if err != nil {
		t.Fatal(err)
	}
}
