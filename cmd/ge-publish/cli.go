package main

import (
	"log"
	"os"

	"github.com/grassrootseconomics/ge-publish/internal/command"
	"github.com/urfave/cli/v2"
)

var (
	version = "dev"
)

func main() {
	command := command.NewCommandContainer(command.CommandOpts{})

	app := &cli.App{
		Name:    "ge-publish",
		Version: version,
		Usage:   "CLI tool to publish GE related smart contracts",
		Commands: []*cli.Command{
			{
				Name:    "publish",
				Aliases: []string{"p"},
				Usage:   "Publish a smart contract",
				Subcommands: []*cli.Command{
					command.RegisterSwapPoolCommand(),
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "private-key",
				Aliases: []string{"y"},
				Usage:   "Private key hex",
			},
			&cli.StringFlag{
				Name:    "rpc",
				Aliases: []string{"p"},
				Usage:   "RPC provider",
			},
			&cli.Uint64Flag{
				Name:  "gas-limit",
				Value: 10_000_000,
				Usage: "Gas limit",
			},
			&cli.Uint64Flag{
				Name:  "gas-fee-cap",
				Value: 10_000_000_000,
				Usage: "Gas fee cap",
			},
			&cli.Uint64Flag{
				Name:  "gas-tip-cap",
				Value: 5,
				Usage: "Gas tip cap",
			},
			&cli.BoolFlag{
				Name:  "testnet",
				Value: false,
				Usage: "Testnet",
			},
			&cli.BoolFlag{
				Name:  "raw-tx",
				Value: false,
				Usage: "Dump raw tx hex",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
