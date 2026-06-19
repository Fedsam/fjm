package install

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/Fedsam/fjm/internal/platform"
)

func Extract(currentOS platform.OS, src, destDir string) error {
	archiveType, err := detectArchiveType(src)
	if err != nil {
		return err
	}

	if archiveType == "zip" {
		if err := ExtractZip(src, destDir); err != nil {
			return err
		}
	} else {
		if err := ExtractTarGz(currentOS, src, destDir); err != nil {
			return err
		}
	}
	return nil
}

func ExtractZip(zipPath, destDir string) error {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		stripedPath := strings.SplitN(filepath.ToSlash(f.Name), "/", 2)
		if len(stripedPath) < 2 {
			continue
		}

		if stripedPath[1] == "" {
			continue
		}

		fpath := filepath.Join(destDir, stripedPath[1])

		if !strings.HasPrefix(filepath.Clean(fpath), filepath.Clean(destDir)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", fpath)
		}

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(fpath, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		if err := extractFile(rc, fpath, f.Mode()); err != nil {
			rc.Close()
			return err
		}
		rc.Close()
	}

	if len(r.File) == 0 {
		return fmt.Errorf("empty archive")
	}

	return nil
}

func ExtractTarGz(currentOS platform.OS, src, destDir string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	gzr, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	isMac := currentOS.String() == platform.Macos.String()

	for {
		headers, err := tr.Next()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return err
		}

		var lvlToStrip int
		if isMac {
			lvlToStrip = 3
		} else {
			lvlToStrip = 2
		}

		stripedPath := strings.SplitN(headers.Name, "/", lvlToStrip)
		if len(stripedPath) < lvlToStrip {
			continue
		}

		if isMac && !(stripedPath[1] == "Contents" && strings.HasPrefix(stripedPath[2], "Home")) {
			continue
		}

		path := filepath.Join(destDir, strings.TrimPrefix(stripedPath[len(stripedPath)-1], "Home/"))

		switch headers.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(path, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
				return err
			}
			if err = extractFile(tr, path, headers.FileInfo().Mode()); err != nil {
				return err
			}
		}
	}
	return nil
}

func detectArchiveType(src string) (string, error) {
	f, err := os.Open(src)
	if err != nil {
		return "", err
	}
	defer f.Close()

	magic := make([]byte, 4)
	n, err := io.ReadFull(f, magic)
	if err != nil && !errors.Is(err, io.ErrUnexpectedEOF) {
		return "", err
	}

	magic = magic[:n]

	if bytes.HasPrefix(magic, []byte{'P', 'K'}) {
		return "zip", nil
	}

	if bytes.HasPrefix(magic, []byte{0x1F, 0x8B}) {
		return "tar.gz", nil
	}

	return "", fmt.Errorf("unknown archive format")
}

func extractFile(r io.Reader, destPath string, mode os.FileMode) error {
	out, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, r)
	return err
}
