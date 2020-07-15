package resources

import "os/exec"

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
