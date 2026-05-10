package shell

import (
	"fmt"
	"strings"
)

type Powershell struct{}

func (s *Powershell) RenderEnv(javaHome string) string {
	envs := fmt.Sprintf(`
$env:JAVA_HOME = "%s"
$javaBin  = Join-Path $env:JAVA_HOME "bin"
$env:PATH = "$javaBin;$env:PATH"
		`,
		javaHome,
	)
	return strings.TrimSpace(envs)
}

func (s *Powershell) SetEnvVar(name, value string) string {
	return fmt.Sprintf(`$env:%s = "%s"`, name, value)
}
