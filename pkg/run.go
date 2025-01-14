package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func StartOpGeth() error {
	homeDir, _ := os.UserHomeDir()
	gethDir := filepath.Join(homeDir, "op-geth")
	datadir := filepath.Join(gethDir, "datadir")
	genesis := filepath.Join(gethDir, "genesis.json")
	jwt := filepath.Join(gethDir, "jwt.txt")

	// 1. init
	initCmd := exec.Command("./build/bin/geth", "init",
		"--state.scheme=hash",
		"--datadir="+datadir,
		genesis,
	)
	initCmd.Dir = gethDir
	initCmd.Stdout = os.Stdout
	initCmd.Stderr = os.Stderr
	if err := initCmd.Run(); err != nil {
		return fmt.Errorf("failed to init op-geth: %w", err)
	}

	// 2. start
	gethCmd := exec.Command("./build/bin/geth",
		"--datadir="+datadir,
		"--http",
		"--http.corsdomain=*",
		"--http.vhosts=*",
		"--http.addr=0.0.0.0",
		"--http.api=web3,debug,eth,txpool,net,engine",
		"--ws",
		"--ws.addr=0.0.0.0",
		"--ws.port=8546",
		"--ws.origins=*",
		"--ws.api=debug,eth,txpool,net,engine",
		"--syncmode=full",
		"--gcmode=archive",
		"--nodiscover",
		"--maxpeers=0",
		"--networkid=42069",
		"--authrpc.vhosts=*",
		"--authrpc.addr=0.0.0.0",
		"--authrpc.port=8551",
		"--authrpc.jwtsecret="+jwt,
		"--rollup.disabletxpoolgossip=true",
	)
	gethCmd.Dir = gethDir
	gethCmd.Stdout = os.Stdout
	gethCmd.Stderr = os.Stderr
	return gethCmd.Run()
}
