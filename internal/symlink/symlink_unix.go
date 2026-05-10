//go:build !windows

package symlink

import "os"

func CreateSymlink(target, link string) error {
	err := os.Symlink(target, link)
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
