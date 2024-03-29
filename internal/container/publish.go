package container

import (
	"fmt"
	"strings"

	"github.com/grassrootseconomics/celoutils/v2"
	"github.com/grassrootseconomics/ge-publish/pkg/decimalquote"
	"github.com/grassrootseconomics/ge-publish/pkg/limiter"
	"github.com/grassrootseconomics/ge-publish/pkg/limiterindex"
	"github.com/grassrootseconomics/ge-publish/pkg/priceindexquoter"
	"github.com/grassrootseconomics/ge-publish/pkg/swappool"
	"github.com/grassrootseconomics/ge-publish/pkg/tokenindex"
	"github.com/urfave/cli/v2"
)

func (c *Container) RegisterPublishCommands() []*cli.Command {
	return []*cli.Command{
		c.decimalQuote(),
		c.limiter(),
		c.limiterIndex(),
		c.priceIndexQuoter(),
		c.swapPool(),
		c.tokenIndex(),
	}
}

func (c *Container) decimalQuote() *cli.Command {
	return &cli.Command{
		Name:  "decimal-quote",
		Usage: "Publish the decimal quote smart contract",
		Action: func(cCtx *cli.Context) error {
			contract := decimalquote.NewDecimalQuoteContract()
			bytecode, err := contract.Bytecode(decimalquote.DecimalQuoteConstructorArgs{})
			if err != nil {
				return err
			}

			resp, err := c.SendContractPublishTx(cCtx, bytecode, contract.GasLimit())
			if err != nil {
				return err
			}

			fmt.Println(resp.TxHash.String())

			return nil
		},
	}
}

func (c *Container) limiter() *cli.Command {
	return &cli.Command{
		Name:  "limiter",
		Usage: "Publish the limiter smart contract",
		Action: func(cCtx *cli.Context) error {
			contract := limiter.NewLimiterContract()
			bytecode, err := contract.Bytecode(limiter.LimiterConstructorArgs{})
			if err != nil {
				return err
			}

			resp, err := c.SendContractPublishTx(cCtx, bytecode, contract.GasLimit())
			if err != nil {
				return err
			}

			fmt.Println(resp.TxHash.String())

			return nil
		},
	}
}

func (c *Container) limiterIndex() *cli.Command {
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

			resp, err := c.SendContractPublishTx(cCtx, bytecode, contract.GasLimit())
			if err != nil {
				return err
			}

			fmt.Println(resp.TxHash.String())

			return nil
		},
	}
}

func (c *Container) priceIndexQuoter() *cli.Command {
	return &cli.Command{
		Name:  "price-index-quoter",
		Usage: "Publish the price index quoter smart contract",
		Action: func(cCtx *cli.Context) error {
			contract := priceindexquoter.NewPriceIndexQuoterContract()
			bytecode, err := contract.Bytecode(priceindexquoter.PriceIndexQuoterConstructorArgs{})
			if err != nil {
				return err
			}

			resp, err := c.SendContractPublishTx(cCtx, bytecode, contract.GasLimit())
			if err != nil {
				return err
			}

			fmt.Println(resp.TxHash.String())

			return nil
		},
	}
}

func (c *Container) swapPool() *cli.Command {
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

			resp, err := c.SendContractPublishTx(cCtx, bytecode, contract.GasLimit())
			if err != nil {
				return err
			}

			fmt.Println(resp.TxHash.String())

			return nil
		},
	}
}

func (c *Container) tokenIndex() *cli.Command {
	return &cli.Command{
		Name:    "token-index",
		Aliases: []string{"token-registry"},
		Usage:   "Publish the ERC20 unique token index smart contract",
		Action: func(cCtx *cli.Context) error {
			contract := tokenindex.NewTokenIndexContract()
			bytecode, err := contract.Bytecode(tokenindex.TokenIndexConstructorArgs{})
			if err != nil {
				return err
			}

			resp, err := c.SendContractPublishTx(cCtx, bytecode, contract.GasLimit())
			if err != nil {
				return err
			}

			fmt.Println(resp.TxHash.String())

			return nil
		},
	}
}
