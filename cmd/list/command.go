package list

import (
	"fmt"

	"github.com/Fedsam/fjm/internal/install"
	"github.com/Fedsam/fjm/internal/manager"
	"github.com/spf13/cobra"
)

func NewCommand(m *manager.FJM) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list [version]",
		Short: "List installed jdk versions",
		RunE: func(cmd *cobra.Command, args []string) error {
			versions, err := install.ListVersions(install.ASC)
			if err != nil {
				return err
			}
			for _, v := range versions {
				fmt.Println(v)
			}
			return nil
		},
	}
	return cmd
}
