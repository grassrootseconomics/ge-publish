package publish

import (
	"fmt"
	"strings"

	"github.com/grassrootseconomics/celoutils/v2"
	"github.com/grassrootseconomics/ge-publish/internal/provider"
	"github.com/grassrootseconomics/ge-publish/pkg/swappool"
	"github.com/urfave/cli/v2"
)

func (c *Command) RegisterSwapPoolCommand() *cli.Command {
	return &cli.Command{
		Name:    "swap-pool",
		Aliases: []string{"pool"},
		Usage:   "Publish the ERC20 swap pool smart contract",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Usage:    "Swap pool name",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "symbol",
				Usage:    "Swap pool symbol",
				Required: true,
				Action: func(ctx *cli.Context, s string) error {
					if len(s) < 3 || len(s) > 10 {
						return fmt.Errorf("flag symbol %s length out of range[3-10]", s)
					}
					return nil
				},
			},
			&cli.UintFlag{
				Name:     "decimals",
				Usage:    "Swap pool decimals",
				Required: true,
				Action: func(ctx *cli.Context, u uint) error {
					if u == 0 || u > 18 {
						return fmt.Errorf("flag decimals value %d out of range[1-18]", u)
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:     "token-registry",
				Usage:    "Swap pool token registry",
				Value:    "0x0000000000000000000000000000",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "token-limiter",
				Value:    "0x0000000000000000000000000000",
				Usage:    "Swap pool token limiter",
				Required: false,
			},
		},
		Action: func(cCtx *cli.Context) error {
			contract := swappool.NewSwapPoolContract()
			bytecode, err := contract.Bytecode(
				swappool.SwapPoolConstructorArgs{
					Name:          cCtx.String("name"),
					Symbol:        strings.ToUpper(cCtx.String("symbol")),
					Decimals:      uint8(cCtx.Uint("decimals")),
					TokenRegistry: celoutils.HexToAddress(cCtx.String("token-registry")),
					TokenLimiter:  celoutils.HexToAddress(cCtx.String("token-limiter")),
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
