//go:build windows

package paths

import (
	"os"
	"path/filepath"
)

func dataHomeDir() string {
	return filepath.Join(os.Getenv("LOCALAPPDATA"), "fjm")
}
