package cert

import (
	"github.com/Fedsam/fjm/internal/cert"
	"github.com/spf13/cobra"
)

func newAddCommand() *cobra.Command {
	var name string
	cmd := &cobra.Command{
		Use:   "add [url]",
		Short: "Add a certificate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			certUrl := args[0]
			alias, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}

			if err := cert.Install(alias, certUrl); err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&name, "name", "n", "", "Certificate name")
	cmd.MarkFlagRequired("name")
	return cmd
}
