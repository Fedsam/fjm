package cert

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Fedsam/fjm/internal/constants"
	"github.com/Fedsam/fjm/internal/install"
)

func Install(alias, hostname string) error {
	jdkPath := os.Getenv(constants.JAVA_HOME)
	filename, err := downloadCertificate(jdkPath, alias, hostname)
	if err != nil {
		return err
	}

	version, err := install.GetCurrentVersion()
	if err != nil {
		return err
	}
	major, err := strconv.ParseInt(strings.SplitN(version, ".", 2)[0], 10, 0)
	if err != nil {
		return err
	}

	keytoolPath, keyStorePath := resolveKeytoolPaths(jdkPath, int(major))
	if err = remove(hostname, keytoolPath, keyStorePath); err != nil {
		return err
	}
	if err = add(filename, alias, keytoolPath, keyStorePath); err != nil {
		return err
	}
	fmt.Printf("certificate %s installed successfully!\n", filename)
	return nil
}
