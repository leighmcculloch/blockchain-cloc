package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Blockchain struct {
	Name  string
	Repos []string
}

var blockchains = []Blockchain{
	{
		Name: "Stellar (C++)",
		Repos: []string{
			"https://github.com/stellar/stellar-core",
			"https://github.com/xdrpp/xdrpp",
		},
	},
	{
		Name: "Bitcoin (C++)",
		Repos: []string{
			"https://github.com/bitcoin/bitcoin",
		},
	},
	{
		Name: "Polkadot (Rust)",
		Repos: []string{
			"https://github.com/paritytech/polkadot",
			"https://github.com/paritytech/substrate",
		},
	},
	{
		Name: "Cosmos (Go)",
		Repos: []string{
			"https://github.com/cosmos/gaia",
			"https://github.com/cosmos/cosmos-sdk",
			"https://github.com/cosmos/ibc-go",
			"https://github.com/cosmos/go-bip39",
			"https://github.com/cosmos/iavl",
			"https://github.com/cosmos/ledger-cosmos-go",
			"https://github.com/cosmos/ledger-go",
			"https://github.com/tendermint/tendermint",
			"https://github.com/tendermint/tm-db",
			"https://github.com/tendermint/btcd",
			"https://github.com/tendermint/crypto",
			"https://github.com/tendermint/go-amino",
		},
	},
	{
		Name: "Ethereum (Go)",
		Repos: []string{
			"https://github.com/ethereum/go-ethereum",
		},
	},
	{
		Name: "Solana (Rust)",
		Repos: []string{
			"https://github.com/solana-labs/solana",
		},
	},
}

func main() {
	err := run()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	for _, b := range blockchains {
		wd, err := ioutil.TempDir("", "")
		if err != nil {
			return fmt.Errorf("creating temp dir: %w", err)
		}

		fmt.Printf("Blockchain: %s\n", b.Name)
		for _, r := range b.Repos {
			c := exec.Command("git", "clone", "--depth=1", r)
			c.Dir = wd
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			err = c.Run()
			if err != nil {
				return fmt.Errorf("exec 'git clone %s': %w", r, err)
			}
		}

		subFiles, err := os.ReadDir(wd)
		if err != nil {
			return fmt.Errorf("reading temp dir: %w", err)
		}
		subDirs := []string{}
		for _, f := range subFiles {
			if f.IsDir() {
				subDirs = append(subDirs, f.Name())
			}
		}

		c := exec.Command("scc", append([]string{"--avg-wage=150000"}, subDirs...)...)
		c.Dir = wd
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		err = c.Run()
		if err != nil {
			return fmt.Errorf("exec 'scc %s': %w", strings.Join(subDirs, " "), err)
		}
	}
	return nil
}
