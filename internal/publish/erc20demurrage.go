package publish

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/grassrootseconomics/celoutils/v2"
	"github.com/grassrootseconomics/ge-publish/internal/provider"
	"github.com/grassrootseconomics/ge-publish/pkg/erc20demurrage"
	"github.com/urfave/cli/v2"
)

func (c *Command) RegisterERC20DemurrageCommand() *cli.Command {
	return &cli.Command{
		Name:    "erc20-demurrage",
		Aliases: []string{"erc20", "det", "voucher", "token"},
		Usage:   "Publish the ERC20 (demurrage) smart contract",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Usage:    "Token name",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "symbol",
				Usage:    "Token symbol",
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
				Usage:    "Token decimals",
				Value:    6,
				Required: false,
				Action: func(ctx *cli.Context, u uint) error {
					if u == 0 || u > 18 {
						return fmt.Errorf("flag decimals value %d out of range[1-18]", u)
					}
					return nil
				},
			},
			&cli.Uint64Flag{
				Name:     "demurrage-level",
				Aliases:  []string{"decay-level"},
				Usage:    "Level of decay per minute",
				Value:    20000,
				Required: false,
			},
			&cli.Uint64Flag{
				Name:     "redistribution-period",
				Aliases:  []string{"period-minutes"},
				Usage:    "Number of minutes between each time the demurraged value can be withdrawn to the Sink Account",
				Value:    43200,
				Required: false,
			},
			&cli.StringFlag{
				Name:     "sink-address",
				Aliases:  []string{"community-fund"},
				Usage:    "The initial Sink Address",
				Required: true,
			},
		},
		Action: func(cCtx *cli.Context) error {
			contract := erc20demurrage.NewERC20DemurrageContract()
			bytecode, err := contract.Bytecode(
				erc20demurrage.ERC20DemurrageConstructorArgs{
					Name:               cCtx.String("name"),
					Symbol:             strings.ToUpper(cCtx.String("symbol")),
					Decimals:           uint8(cCtx.Uint("decimals")),
					DecayLevel:         big.NewInt(int64(cCtx.Uint64("demurrage-level"))),
					PeriodMinutes:      big.NewInt(int64(cCtx.Uint64("redistribution-period"))),
					DefaultSinkAddress: celoutils.HexToAddress(cCtx.String("sink-address")),
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
