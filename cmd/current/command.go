package current

import (
	"fmt"

	"github.com/Fedsam/fjm/internal/install"
	"github.com/Fedsam/fjm/internal/manager"
	"github.com/spf13/cobra"
)

func NewCommand(m *manager.FJM) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "current [version]",
		Short: "Make specified jdk version default",
		RunE: func(cmd *cobra.Command, args []string) error {
			version, err := install.GetCurrentVersion()
			if err != nil {
				return err
			}
			fmt.Println(version)
			return nil
		},
	}
	return cmd
}
