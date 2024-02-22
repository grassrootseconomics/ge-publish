package publish

import (
	"fmt"

	"github.com/grassrootseconomics/celoutils/v2"
	"github.com/grassrootseconomics/ge-publish/internal/provider"
	"github.com/grassrootseconomics/ge-publish/pkg/limiterindex"
	"github.com/urfave/cli/v2"
)

func (c *Command) RegisterLimiterIndexCommand() *cli.Command {
	return &cli.Command{
		Name:  "limiter-index",
		Usage: "Publish the limiter index smart contract",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "holder",
				Usage:    "Holder",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "limiter-address",
				Usage:    "Existing limiter smart contract address",
				Required: true,
			},
		},
		Action: func(cCtx *cli.Context) error {
			contract := limiterindex.NewLimiterIndexContract()
			bytecode, err := contract.Bytecode(
				limiterindex.LimiterIndexConstructorArgs{
					Holder:         celoutils.HexToAddress(cCtx.String("holder")),
					LimiterAddress: celoutils.HexToAddress(cCtx.String("limiter-address")),
				},
			)
			if err != nil {
				return err
			}

			resp, err := provider.SendContractPublishTx(cCtx, bytecode, contract.GasLimit())
			if err != nil {
				return err
			}

			fmt.Println(resp.TxHash.String())

			return nil
		},
	}
}
