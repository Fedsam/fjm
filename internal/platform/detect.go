package platform

import (
	"runtime"
)

func DetectOS() OS {
	switch runtime.GOOS {
	case "windows":
		return Windows
	case "darwin":
		return Macos
	default:
		return Linux
	}
}

func DetectArch() Arch {
	switch runtime.GOARCH {
	case "arm64":
		return ARM64
	case "386":
		return X32
	default: // "amd64"
		return X64
	}
}
