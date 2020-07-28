package utils

import (
	"crypto/rand"
	"encoding/base32"
	"errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
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
	if err := os.MkdirAll(filepath.Dir(p), 0777); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func WriteFileString(path, content string) error {
	if _, err := os.Stat(path); err == nil {
		return errors.New("file already exists")
	}

	return AppendFileString(path, content)
}

func AppendFileString(path, content string) error {
	var f *os.File = nil
	finfo, err := os.Stat(path)
	if err == nil {
		f, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, finfo.Mode())
		if err != nil {
			return err
		}
	} else if os.IsNotExist(err) {
		f, err = CreateFileRecursive(path)
		if err != nil {
			return err
		}
	} else {
		return err
	}

	defer f.Close()
	if _, err := f.WriteString(content); err != nil {
		return err
	}

	return f.Sync()
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

var rnd uint32
var rndmu sync.Mutex

func reseed() uint32 {
	return uint32(time.Now().UnixNano() + int64(os.Getpid()))
}

func nextRandom() string {
	rndmu.Lock()
	r := rnd
	if r == 0 {
		r = reseed()
	}
	r = r*1664525 + 1013904223 // constants from Numerical Recipes
	rnd = r
	rndmu.Unlock()
	return strconv.Itoa(int(1e9 + r%1e9))[1:]
}

func GenerateTempFileName(dir, pattern string) (string, error) {
	var prefix, suffix string
	if pos := strings.LastIndex(pattern, "*"); pos != -1 {
		prefix, suffix = pattern[:pos], pattern[pos+1:]
	} else {
		prefix = pattern
	}
	nconflict := 0

	for i := 0; i < 10000; i++ {
		name := filepath.Join(dir, prefix+nextRandom()+suffix)

		f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)
		if os.IsExist(err) {
			if nconflict++; nconflict > 10 {
				rndmu.Lock()
				rnd = reseed()
				rndmu.Unlock()
			}
			continue
		}
		defer f.Close()
		return name, nil
	}

	return "", errors.New("max retry exceeded")
}
