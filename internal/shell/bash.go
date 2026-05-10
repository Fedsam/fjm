package shell

import (
	"fmt"
	"strings"
)

type Bash struct{}

func (s *Bash) RenderEnv(javaHome string) string {
	envs := fmt.Sprintf(`
export JAVA_HOME="%s"
export PATH="$JAVA_HOME/bin:$PATH"
		`,
		javaHome,
	)
	return strings.TrimSpace(envs)
}

func (s *Bash) SetEnvVar(name, value string) string {
	return fmt.Sprintf(`export %s="%s"`, name, value)
}
