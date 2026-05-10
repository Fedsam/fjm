package paths

import (
	"os"
	"path/filepath"
)

const (
	MultishellDir = `\fjm_multishells`
)

func VersionsDir() string {
	return filepath.Join(dataHomeDir(), "versions")
}

func AliasesDir() string {
	return filepath.Join(dataHomeDir(), "aliases")
}

func CertsDir() string {
	return filepath.Join(dataHomeDir(), "certs")
}

func dataHomeDir() string {
	state := os.Getenv("XDG_DATA_HOME")
	if state == "" {
		home, _ := os.UserHomeDir()
		state = filepath.Join(home, ".local", "share", "fjm")
	}
	return state
}
