package env

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Fedsam/fjm/internal/constants"
	"github.com/Fedsam/fjm/internal/install"
	"github.com/Fedsam/fjm/internal/manager"
	"github.com/Fedsam/fjm/internal/paths"
	"github.com/Fedsam/fjm/internal/symlink"
	"github.com/spf13/cobra"
)

func NewCommand(m *manager.FJM) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "env",
		Short: "Print shell environment for current Java",
		RunE: func(cmd *cobra.Command, args []string) error {
			cacheDir, err := os.UserCacheDir()
			if err != nil {
				return err
			}
			multishellDir := filepath.Join(cacheDir, constants.FJM_MULTISHELL_CACHE, generateSymLinkPath())

			err = os.MkdirAll(multishellDir, 0755)
			if err != nil {
				fmt.Println("Error creating folder:", err)
				return err
			}

			jdkHome := filepath.Join(multishellDir, "jdk")

			script := m.Shell.RenderEnv(jdkHome)
			_, err = fmt.Fprintln(cmd.OutOrStdout(), script)
			if err != nil {
				return err
			}

			multishell := m.Shell.SetEnvVar(constants.FJM_MULTISHELL_PATH, multishellDir)
			_, err = fmt.Fprintln(cmd.OutOrStdout(), multishell)
			if err != nil {
				return err
			}

			if err = setDefaultValue(multishellDir); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

func generateSymLinkPath() string {
	return fmt.Sprintf("%d_%d", os.Getpid(), time.Now().UnixMilli())
}

func setDefaultValue(multishellPath string) error {
	jdkDir := filepath.Join(multishellPath, "jdk")

	versions, err := install.ListVersions(install.ASC)
	if err != nil {
		return err
	}

	if len(versions) == 0 {
		return nil // silence no versions installed
	}

	target := filepath.Join(paths.VersionsDir(), versions[len(versions)-1].String())
	if err := symlink.CreateSymlink(target, jdkDir); err != nil {
		return err
	}
	return nil
}
