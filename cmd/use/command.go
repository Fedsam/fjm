package use

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Fedsam/fjm/internal/constants"
	"github.com/Fedsam/fjm/internal/install"
	"github.com/Fedsam/fjm/internal/manager"
	"github.com/Fedsam/fjm/internal/paths"
	"github.com/Fedsam/fjm/internal/symlink"
	"github.com/hashicorp/go-version"
	"github.com/spf13/cobra"
)

func NewCommand(m *manager.FJM) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "use [version]",
		Short: "Make specified jdk version default",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var v *version.Version

			versions, err := install.ListVersions(install.DESC)
			if err != nil {
				return err
			}

			if len(args) == 0 {
				return fmt.Errorf("Missing version number")
			}

			v, _ = version.NewVersion(args[0])
			isExactMatch := len(strings.SplitN(args[0], ".", 3)) == 1

			matchIndex := sort.Search(len(versions), func(i int) bool {
				if isExactMatch {
					return versions[i].Equal(v)
				}
				return versions[i].Segments()[0] == v.Segments()[0]
			})

			if matchIndex >= len(versions) {
				matchIndex = len(versions) - 1
			}

			multishellPath := os.Getenv(constants.FJM_MULTISHELL_PATH)
			jdkDir := filepath.Join(multishellPath, "jdk")
			target := filepath.Join(paths.VersionsDir(), versions[matchIndex].String())

			if err := symlink.DeleteSymlink(jdkDir); err != nil {
				return err
			}
			if err := symlink.CreateSymlink(target, jdkDir); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}
