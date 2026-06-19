package install

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Fedsam/fjm/internal/paths"
	"github.com/Fedsam/fjm/internal/platform"
)

func Install(p *platform.Platform, v int8) error {
	versionsDir := paths.VersionsDir()

	tmpDir, err := os.MkdirTemp(versionsDir, "download-*")
	if err != nil {
		return err
	}

	fmt.Printf("downloading jdk %d...\n", v)
	tmpFile, err := Download(buildUrl(p, v), tmpDir)
	if err != nil {
		_ = os.RemoveAll(tmpDir)
		return fmt.Errorf("installing jdk %d, cause: %w", v, err)
	}

	if err = Extract(p.OS, tmpFile, tmpDir); err != nil {
		_ = os.RemoveAll(tmpDir)
		return err
	}

	if err = os.Remove(tmpFile); err != nil {
		return err
	}

	release, err := ParseRelease(filepath.Join(tmpDir, "release"))
	if err != nil {
		return err
	}

	installedVersion := release["JAVA_VERSION"]

	finalPath := filepath.Join(versionsDir, installedVersion)
	if err = os.Rename(tmpDir, finalPath); err != nil {
		_ = os.RemoveAll(tmpDir)
		return fmt.Errorf("install jdk %s: %w", installedVersion, err)
	}

	fmt.Printf("jdk %s installed successfully!\n", installedVersion)
	return nil
}

func buildUrl(p *platform.Platform, version int8) string {
	const jdkUrl = "https://api.adoptium.net/v3/binary/latest/%d/ga/%s/%s/jdk/hotspot/normal/eclipse"
	return fmt.Sprintf(jdkUrl, version, p.OS.String(), p.Arch.String())
}
