package install

import (
	"errors"
	"io"
	"net/http"
	"os"
	"time"
)

// Download downloads a file from url and saves it as a temp file in destDir.
// Returns the path to the downloaded file.
func Download(url, destDir string) (string, error) {
	client := &http.Client{Timeout: 5 * time.Minute}

	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", errors.New(string(body))
	}

	f, err := os.CreateTemp(destDir, "downloadf-*")
	if err != nil {
		return "", err
	}
	defer f.Close()

	if _, err := io.Copy(f, resp.Body); err != nil {
		os.Remove(f.Name())
		return "", err
	}

	return f.Name(), nil
}
