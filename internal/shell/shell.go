package shell

import (
	"os"
	"path/filepath"
)

type Shell interface {
	RenderEnv(javaHome string) string
	SetEnvVar(name, value string) string
}

func Detect() Shell {
	shell := os.Getenv("SHELL")
	switch filepath.Base(shell) {
	case "bash":
		return &Bash{}
	case "zsh":
		return &Zsh{}
	default:
		return &Powershell{}
	}
}
