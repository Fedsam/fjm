//go:build windows

package symlink

import (
	"os"
	"os/exec"
)

func CreateSymlink(target, link string) error {
	cmd := exec.Command("cmd", "/c", "mklink", "/J", link, target)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func DeleteSymlink(link string) error {
	if _, err := os.Lstat(link); err == nil {
		if err := os.Remove(link); err != nil {
			return err
		}
	}
	return nil
}
