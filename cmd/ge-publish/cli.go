package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/grassrootseconomics/ge-publish/internal/command"
	"github.com/kamikazechaser/common/logg"
	"github.com/urfave/cli/v2"
)

var (
	version = "dev"
)

func main() {
	loggOpts := logg.LoggOpts{
		FormatType: logg.Human,
		LogLevel:   slog.LevelInfo,
	}

	if os.Getenv("DEBUG") != "" {
		loggOpts.LogLevel = slog.LevelDebug
	}

	baseLogger := logg.NewLogg(loggOpts)
	commands := command.NewCommands()

	app := &cli.App{
		Name:    "ge-publish",
		Version: version,
		Usage:   "CLI tool to publish GE related smart contracts",
		Before: func(cCtx *cli.Context) error {

			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "publish",
				Aliases: []string{"p"},
				Usage:   "Publish a smart contract",
				Subcommands: []*cli.Command{
					commands.RegisterSwapPoolCommand(baseLogger),
					commands.RegisterDecimalQuoteCommand(baseLogger),
					commands.RegisterPriceIndexQuoteCommand(baseLogger),
					commands.RegisterLimiterCommand(baseLogger),
					commands.RegisterLimiterIndexCommand(baseLogger),
					commands.RegisterTokenIndexCommand(baseLogger),
				},
			},
			commands.RegisterWritePrivateKeyCommand(baseLogger),
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "private-key",
				Aliases:  []string{"y"},
				Usage:    "Private key hex",
				Required: true,
				EnvVars:  []string{"PRIVATE_KEY"},
				FilePath: commands.PrivateKeyFileLocation(baseLogger),
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
		baseLogger.Error("could not run ge-publish", "error", err)
		os.Exit(1)
	}
}
