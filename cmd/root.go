package cmd

import (
	"fmt"
	"os"

	"github.com/Fedsam/fjm/cmd/cert"
	"github.com/Fedsam/fjm/cmd/current"
	"github.com/Fedsam/fjm/cmd/env"
	"github.com/Fedsam/fjm/cmd/install"
	"github.com/Fedsam/fjm/cmd/list"
	"github.com/Fedsam/fjm/cmd/use"
	"github.com/Fedsam/fjm/internal/manager"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "fjm",
	Short:        "A simple cli tool to install/manage jdks and certificates",
	Long:         ``,
	SilenceUsage: true,
}

func Execute() {
	mgr, err := manager.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	rootCmd.AddCommand(env.NewCommand(mgr))
	rootCmd.AddCommand(install.NewCommand(mgr))
	rootCmd.AddCommand(use.NewCommand(mgr))
	rootCmd.AddCommand(list.NewCommand(mgr))
	rootCmd.AddCommand(current.NewCommand(mgr))
	rootCmd.AddCommand(cert.NewCommand(mgr))

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
