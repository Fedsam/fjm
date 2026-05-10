package manager

import (
	"os"

	"github.com/Fedsam/fjm/internal/install"
	"github.com/Fedsam/fjm/internal/paths"
	"github.com/Fedsam/fjm/internal/platform"
	"github.com/Fedsam/fjm/internal/shell"
)

type FJM struct {
	Platform platform.Platform
	Shell    shell.Shell
}

func New() (*FJM, error) {
	versionsDir := paths.VersionsDir()
	certsDir := paths.CertsDir()

	if err := os.MkdirAll(versionsDir, 0755); err != nil {
		return nil, err
	}

	if err := os.MkdirAll(certsDir, 0755); err != nil {
		return nil, err
	}

	return &FJM{Platform: platform.Detect(), Shell: shell.Detect()}, nil
}

func (m *FJM) Install(version int8) error {
	if err := install.Install(&m.Platform, version); err != nil {
		return err
	}
	return nil
}
