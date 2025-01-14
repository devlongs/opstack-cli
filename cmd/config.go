package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config-l2",
	Short: "Generate L2 config files: genesis.json, rollup.json, jwt.txt",
	RunE: func(cmd *cobra.Command, args []string) error {
		homeDir, _ := os.UserHomeDir()
		opNodePath := filepath.Join(homeDir, "optimism", "op-node")

		// 1. Generate the config
		c := exec.Command("go", "run", "cmd/main.go", "genesis", "l2",
			"--deploy-config", "../packages/contracts-bedrock/deploy-config/getting-started.json",
			"--l1-deployments", "../packages/contracts-bedrock/deployments/getting-started/.deploy",
			"--outfile.l2", "genesis.json",
			"--outfile.rollup", "rollup.json",
			"--l1-rpc", os.Getenv("L1_RPC_URL"),
		)
		c.Dir = opNodePath
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			return fmt.Errorf("failed to generate genesis files: %w", err)
		}

		// 2. Generate jwt.txt
		jwtCmd := exec.Command("openssl", "rand", "-hex", "32")
		jwtOut, err := jwtCmd.Output()
		if err != nil {
			return fmt.Errorf("failed to generate JWT: %w", err)
		}
		jwtFile := filepath.Join(opNodePath, "jwt.txt")
		if err := os.WriteFile(jwtFile, jwtOut, 0600); err != nil {
			return fmt.Errorf("failed to write JWT: %w", err)
		}

		fmt.Println("L2 config generated successfully!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
