package resources

import (
	"context"
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
	defer func() {
		os.RemoveAll(path + "/test")
		os.RemoveAll(path + "/output")
	}()
	installer := &RestyInstaller{
		BuildDir:     path + "/test",
		Prefix:       path + "/output",
		BuildOptions: []string{"--with-http_mp4_module"},
	}
	ctx := context.Background()
	err = installer.Install(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMultiStdout(t *testing.T) {
	defer func() {
		os.RemoveAll("log.log")
	}()
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
