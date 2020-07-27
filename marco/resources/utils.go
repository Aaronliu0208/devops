package resources

import (
	"os"
	"os/exec"
)

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
