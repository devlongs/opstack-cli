package pkg

import (
	"fmt"
	"os/exec"
)

func CheckDependencies() error {
	neededBins := []string{"go", "pnpm", "direnv", "forge", "make"}
	for _, bin := range neededBins {
		cmd := exec.Command("which", bin)
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("'%s' is not installed or not in PATH", bin)
		}
	}

	// Todo (Longs): Optionally, check specific versions by invoking bin --version and parsing the output

	return nil
}
