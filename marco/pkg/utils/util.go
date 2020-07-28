package utils

import (
	"crypto/rand"
	"encoding/base32"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var encoding = base32.NewEncoding("abcdefghijklmnopqrstuvwxyz234567")

// NewID returns a random string which can be used as an ID for objects.
func NewID() string {
	buff := make([]byte, 16) // 128 bit random ID.
	if _, err := io.ReadFull(rand.Reader, buff); err != nil {
		panic(err)
	}
	// Avoid the identifier to begin with number and trim padding
	return string(buff[0]%26+'a') + strings.TrimRight(encoding.EncodeToString(buff[1:]), "=")
}

func CreateFileRecursive(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

//CommandExists check cmd exists
func CommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

//IsExecAll To tell if the file is executable by all of the above, again use bitmask 0111 but check if the result equals to 0111
// https://stackoverflow.com/questions/60128401/how-to-check-if-a-file-is-executable-in-go
func IsExecAll(mode os.FileMode) bool {
	return mode&0111 == 0111
}
