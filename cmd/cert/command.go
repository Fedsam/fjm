package cert

import (
	"github.com/Fedsam/fjm/internal/manager"
	"github.com/spf13/cobra"
)

func NewCommand(m *manager.FJM) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cert",
		Aliases: []string{"c"},
		Short:   "Manage jdk certificates",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	cmd.AddCommand(newAddCommand())
	return cmd
}
