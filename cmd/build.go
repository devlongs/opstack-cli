package cmd

import (
	"fmt"

	"github.com/devlongs/opstack-cli/pkg"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the optimism monorepo and op-geth binaries",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := pkg.BuildOptimism(); err != nil {
			return fmt.Errorf("failed to build optimism: %w", err)
		}
		if err := pkg.BuildOpGeth(); err != nil {
			return fmt.Errorf("failed to build op-geth: %w", err)
		}
		fmt.Println("Build complete!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
