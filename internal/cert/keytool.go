package cert

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Fedsam/fjm/internal/constants"
)

func resolveKeytoolPaths(jdkPath string, jdkVersion int) (keytoolPath, keyStorePath string) {
	base := "."
	if jdkVersion <= 8 {
		base = "jre"
	}

	keytoolPath = filepath.Join(jdkPath, base, "bin", "keytool.exe")
	keyStorePath = filepath.Join(jdkPath, base, "lib", "security", "cacerts")
	return
}

func add(certFilename, alias, keytoolPath, keyStorePath string) error {
	cmd := exec.Command(
		keytoolPath,
		"-import",
		"-noprompt",
		"-storepass", "changeit",
		"-file", `.\`+certFilename,
		"-alias", alias,
		"-keystore", keyStorePath,
	)
	cmd.Dir = os.Getenv(constants.JAVA_HOME)

	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

func remove(alias, keytoolPath, keyStorePath string) error {
	cmd := exec.Command(
		keytoolPath,
		"-delete",
		"-alias", alias,
		"-keystore", keyStorePath,
		"-storepass", "changeit",
	)

	// Ignore error: alias may not exist
	if err := cmd.Run(); err != nil {
		return nil
	}
	return nil
}
