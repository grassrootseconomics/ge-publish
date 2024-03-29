package main

import (
	"fmt"
	"log"
	"os"

	"github.com/grassrootseconomics/ge-publish/internal/container"
	"github.com/urfave/cli/v2"
)

var (
	version = "dev"
)

func main() {
	container := container.NewContainer()

	app := &cli.App{
		Name:    "ge-publish",
		Version: version,
		Usage:   "CLI tool to publish GE related smart contracts",
		Commands: []*cli.Command{
			{
				Name:        "publish",
				Aliases:     []string{"p"},
				Usage:       "Publish smart contracts",
				Subcommands: container.RegisterPublishCommands(),
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "private-key",
				Aliases:  []string{"y"},
				Usage:    "Private key hex",
				Required: true,
				EnvVars:  []string{"PRIVATE_KEY"},
			},
			&cli.StringFlag{
				Name:     "rpc",
				Aliases:  []string{"p"},
				Usage:    "RPC provider",
				Required: false,
				Value:    "https://1rpc.io/celo",
				EnvVars:  []string{"RPC_PROVIDER"},
			},
			&cli.Int64Flag{
				Name:    "gas-fee-cap",
				Aliases: []string{"fee", "gas-fee"},
				Value:   10_000_000_000,
				Usage:   "Gas fee cap",
				EnvVars: []string{"GAS_FEE_CAP"},
				Action: func(ctx *cli.Context, i int64) error {
					if i < 5_000_000 {
						return fmt.Errorf("gas fee %d below minimum[5M]", i)
					}
					return nil
				},
			},
			&cli.Int64Flag{
				Name:    "gas-tip-cap",
				Aliases: []string{"tip", "gas-tip"},
				Value:   5,
				Usage:   "Gas tip cap",
				EnvVars: []string{"GAS_TIP_CAP"},
			},
			&cli.BoolFlag{
				Name:    "testnet",
				Value:   false,
				Usage:   "Testnet",
				EnvVars: []string{"TESTNET"},
			},
			&cli.BoolFlag{
				Name:    "vv",
				Aliases: []string{"verbose", "debug"},
				Value:   false,
				Usage:   "Verbose logging",
				EnvVars: []string{"DEBUG"},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
