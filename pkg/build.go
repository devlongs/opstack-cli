package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func CloneRepo(folderName, repoURL, branch string) error {
	// e.g. cd ~ or user-specified
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	clonePath := filepath.Join(homeDir, folderName)

	if _, err := os.Stat(clonePath); !os.IsNotExist(err) {
		fmt.Printf("%s already exists, skipping clone...\n", clonePath)
		return nil
	}

	cmd := exec.Command("git", "clone", repoURL, clonePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	if branch != "" {
		if err := CheckoutBranch(clonePath, branch); err != nil {
			return err
		}
	}

	return nil
}

func CheckoutBranch(clonePath, branch string) error {
	cmd := exec.Command("git", "-C", clonePath, "checkout", branch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func BuildOptimism() error {
	homeDir, _ := os.UserHomeDir()
	optDir := filepath.Join(homeDir, "optimism")

	// 1. pnpm install
	cmd := exec.Command("pnpm", "install")
	cmd.Dir = optDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	// 2. make op-node op-batcher op-proposer
	cmd = exec.Command("make", "op-node", "op-batcher", "op-proposer")
	cmd.Dir = optDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	// 3. pnpm build
	cmd = exec.Command("pnpm", "build")
	cmd.Dir = optDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func BuildOpGeth() error {
	homeDir, _ := os.UserHomeDir()
	gethDir := filepath.Join(homeDir, "op-geth")

	cmd := exec.Command("make", "geth")
	cmd.Dir = gethDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
