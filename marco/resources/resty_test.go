package resources

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
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

	_, exists := installer.CheckExists()
	assert.False(t, exists)
	err = installer.Install(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// check resty exists
	_, exists = installer.CheckExists()
	assert.True(t, exists)

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

func int8ToStr(arr []int8) string {
	b := make([]byte, 0, len(arr))
	for _, v := range arr {
		if v == 0x00 {
			break
		}
		b = append(b, byte(v))
	}
	return string(b)
}
func TestRuntimeInfo(t *testing.T) {
	fmt.Println(runtime.GOOS)
	var uname syscall.Utsname
	if err := syscall.Uname(&uname); err == nil {
		// extract members:
		// type Utsname struct {
		//  Sysname    [65]int8
		//  Nodename   [65]int8
		//  Release    [65]int8
		//  Version    [65]int8
		//  Machine    [65]int8
		//  Domainname [65]int8
		// }

		fmt.Println(int8ToStr(uname.Sysname[:]))
		fmt.Println(int8ToStr(uname.Release[:]))
		fmt.Println(int8ToStr(uname.Version[:]))
		fmt.Println(int8ToStr(uname.Machine[:]))
	}

}
