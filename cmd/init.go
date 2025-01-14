package cmd

import (
	"fmt"

	"github.com/devlongs/opstack-cli/pkg"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize directories, clone repos, and check dependencies",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 1. Check dependencies
		if err := pkg.CheckDependencies(); err != nil {
			return fmt.Errorf("dependency check failed: %w", err)
		}

		// 2. Clone the optimism monorepo
		if err := pkg.CloneRepo("optimism", "https://github.com/ethereum-optimism/optimism.git", "tutorials/chain"); err != nil {
			return fmt.Errorf("failed to clone optimism repo: %w", err)
		}

		// 3. Clone op-geth
		if err := pkg.CloneRepo("op-geth", "https://github.com/ethereum-optimism/op-geth.git", ""); err != nil {
			return fmt.Errorf("failed to clone op-geth: %w", err)
		}

		fmt.Println("Initialization successful.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
