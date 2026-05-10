package install

import (
	"os"
	"path/filepath"
	"testing"
)

const (
	release_content = `
JAVA_VERSION="21.0.3"
IMPLEMENTOR="Eclipse Adoptium"
OS_NAME="Mac OS X"
`
	empty_value_release = `
JAVA_VERSION=
IMPLEMENTOR="Eclipse Adoptium"
OS_NAME="Mac OS X"
`
)

func TestParseRelease(t *testing.T) {
	t.Run("Parse release file", func(t *testing.T) {
		tempFile := createReleaseFile(t, release_content)

		got, err := ParseRelease(tempFile)
		if err != nil {
			t.Fatal(err)
		}
		want := "21.0.3"

		javaVersion := got["JAVA_VERSION"]

		if javaVersion != want {
			t.Errorf("got %q want %q", javaVersion, want)
		}
	})
	t.Run("Should return an error if file does not exist", func(t *testing.T) {
		_, err := ParseRelease("")
		if err == nil {
			t.Errorf("Not existant file should return an error")
		}
	})
	t.Run("Should not return an error if key does not have value", func(t *testing.T) {
		tempFile := createReleaseFile(t, empty_value_release)

		got, err := ParseRelease(tempFile)
		if err != nil {
			t.Fatal(err)
		}

		want := 2

		if len(got) != want {
			t.Errorf("got %d want %d", len(got), want)
		}
	})
}

func createReleaseFile(t *testing.T, content string) string {
	dir := t.TempDir()
	filename := filepath.Join(dir, "release")

	if err := os.WriteFile(filename, []byte(content), 0755); err != nil {
		t.Fatal(err)
	}

	return filename
}
