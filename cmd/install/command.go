package install

import (
	"strconv"

	"github.com/Fedsam/fjm/internal/manager"
	"github.com/spf13/cobra"
)

func NewCommand(m *manager.FJM) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install [version]",
		Short: "Install specified jdk version",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var version int8

			if len(args) == 1 {
				v, _ := strconv.ParseInt(args[0], 10, 8)
				version = int8(v)
			}

			if err := m.Install(version); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}
