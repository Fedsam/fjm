package paths

import (
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
