package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var addressesCmd = &cobra.Command{
	Use:   "addresses",
	Short: "Generate addresses and private keys with the wallets.sh script",
	RunE: func(cmd *cobra.Command, args []string) error {
		homeDir, _ := os.UserHomeDir()
		optDir := filepath.Join(homeDir, "optimism")

		scriptPath := filepath.Join(optDir, "packages", "contracts-bedrock", "scripts", "getting-started", "wallets.sh")
		c := exec.Command("bash", scriptPath)
		c.Dir = optDir
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		if err := c.Run(); err != nil {
			return fmt.Errorf("failed to generate addresses: %w", err)
		}

		fmt.Println("Addresses successfully generated!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addressesCmd)
}
