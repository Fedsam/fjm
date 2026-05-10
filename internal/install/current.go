package install

import (
	"os"
	"path/filepath"

	"github.com/Fedsam/fjm/internal/constants"
)

func GetCurrentVersion() (string, error) {
	currentJdk := os.Getenv(constants.JAVA_HOME)

	res, err := ParseRelease(filepath.Join(currentJdk, "release"))
	if err != nil {
		return "", err
	}
	return res["JAVA_VERSION"], nil
}
