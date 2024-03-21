package main

import (
	"fmt"
	"log"
	"os"

	"github.com/grassrootseconomics/ge-publish/internal/publish"
	"github.com/urfave/cli/v2"
)

var (
	version = "dev"
)

func main() {
	command := publish.NewCommandContainer(publish.CommandOpts{})

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
					command.RegisterDecimalQuoteCommand(),
					command.RegisterPriceIndexQuoterCommand(),
					command.RegisterLimiterCommand(),
					command.RegisterLimiterIndexCommand(),
					command.RegisterTokenIndexCommand(),
					command.RegisterERC20DemurrageCommand(),
				},
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
				Name:    "rpc",
				Aliases: []string{"p"},
				Usage:   "RPC provider",
				EnvVars: []string{"RPC_PROVIDER"},
			},
			&cli.Int64Flag{
				Name:    "gas-fee-cap",
				Value:   10_000_000_000,
				Usage:   "Gas fee cap",
				EnvVars: []string{"GAS_FEE_CAP"},
				Action: func(ctx *cli.Context, i int64) error {
					if i < 5_000_000 {
						return fmt.Errorf("flag gas-fee-cap %d below minimum[5M]", i)
					}
					return nil
				},
			},
			&cli.Int64Flag{
				Name:    "gas-tip-cap",
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
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
