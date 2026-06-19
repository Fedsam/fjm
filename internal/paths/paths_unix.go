//go:build !windows

package paths

import (
	"os"
	"path/filepath"
)

func dataHomeDir() string {
	state := os.Getenv("XDG_DATA_HOME")
	if state == "" {
		home, _ := os.UserHomeDir()
		state = filepath.Join(home, ".local", "share", "fjm")
	}
	return state
}
