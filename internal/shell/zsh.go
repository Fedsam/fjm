package shell

import (
	"fmt"
	"strings"
)

type Zsh struct{}

func (s *Zsh) RenderEnv(javaHome string) string {
	envs := fmt.Sprintf(`
export JAVA_HOME="%s"
export PATH="$JAVA_HOME/bin:$PATH"
`,
		javaHome,
	)
	return strings.TrimSpace(envs)
}

func (s *Zsh) SetEnvVar(name, value string) string {
	return fmt.Sprintf(`export %s="%s"`, name, value)
}
