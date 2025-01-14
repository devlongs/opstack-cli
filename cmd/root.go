package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "opstack-cli",
	Short: "A CLI tool to spin up an OP Stack–based rollup",
	Long: `opstack-cli allows you to quickly deploy, configure, and run an OP Stack 
rollup using locally built binaries. This tool wraps standard commands from 
the official Optimism monorepo setup guide.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use one of the subcommands. For help: opstack-cli --help")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Todo (longs): define any persistent flags, or subcommands
}
