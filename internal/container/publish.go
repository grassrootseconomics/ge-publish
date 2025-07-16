package container

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/grassrootseconomics/ethutils"
	"github.com/grassrootseconomics/ge-publish/pkg/contract"
	"github.com/grassrootseconomics/ge-publish/pkg/util"
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
		c.erc20Demurrage(),
		c.periodSimple(),
		c.ethFaucet(),
		c.accountsIndex(),
		c.erc20(),
		c.contractsRegistry(),
		c.custodialRegistrationProxy(),
	}
}

func (c *Container) decimalQuote() *cli.Command {
	return &cli.Command{
		Name:  "decimal-quote",
		Usage: "Publish the decimal quote smart contract",
		Action: func(cCtx *cli.Context) error {
			contract := contract.NewDecimalQuote()
			bytecode, err := contract.Bytecode()
			if err != nil {
				return err
			}
			c.logInitStage(contract)

			resp, err := c.SendContractPublishTx(cCtx, bytecode, contract.MaxGasLimit())
			if err != nil {
				return err
			}
			c.logPublishedStage(contract, resp)

			return nil
		},
	}
}

func (c *Container) limiter() *cli.Command {
	return &cli.Command{
		Name:  "limiter",
		Usage: "Publish the limiter smart contract",
		Action: func(cCtx *cli.Context) error {
			contract := contract.NewLimiter()
			bytecode, err := contract.Bytecode()
			if err != nil {
				return err
			}
			c.logInitStage(contract)

			resp, err := c.SendContractPublishTx(cCtx, bytecode, contract.MaxGasLimit())
			if err != nil {
				return err
			}
			c.logPublishedStage(contract, resp)

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
			contract := contract.NewLimiterIndex(contract.LimiterIndexConstructorArgs{
				Holder:         ethutils.HexToAddress(cCtx.String("holder")),
				LimiterAddress: ethutils.HexToAddress(cCtx.String("limiter-address")),
			})
			bytecode, err := contract.Bytecode()
			if err != nil {
				return err
			}
			c.logInitStage(contract)

			resp, err := c.SendContractPublishTx(cCtx, bytecode, contract.MaxGasLimit())
			if err != nil {
				return err
			}
			c.logPublishedStage(contract, resp)

			return nil
		},
	}
}

func (c *Container) priceIndexQuoter() *cli.Command {
	return &cli.Command{
		Name:  "price-index-quoter",
		Usage: "Publish the price index quoter smart contract",
		Action: func(cCtx *cli.Context) error {
			contract := contract.NewPriceIndexQuoter()
			bytecode, err := contract.Bytecode()
			if err != nil {
				return err
			}
			c.logInitStage(contract)

			resp, err := c.SendContractPublishTx(cCtx, bytecode, contract.MaxGasLimit())
			if err != nil {
				return err
			}
			c.logPublishedStage(contract, resp)

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
			contract := contract.NewSwapPool(contract.SwapPoolConstructorArgs{
				Name:          cCtx.String("name"),
				Symbol:        strings.ToUpper(cCtx.String("symbol")),
				Decimals:      uint8(cCtx.Uint("decimals")),
				TokenRegistry: ethutils.HexToAddress(cCtx.String("token-registry")),
				TokenLimiter:  ethutils.HexToAddress(cCtx.String("token-limiter")),
			})
			bytecode, err := contract.Bytecode()
			if err != nil {
				return err
			}
			c.logInitStage(contract)

			resp, err := c.SendContractPublishTx(cCtx, bytecode, contract.MaxGasLimit())
			if err != nil {
				return err
			}
			c.logPublishedStage(contract, resp)

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
			contract := contract.NewTokenIndex()
			bytecode, err := contract.Bytecode()
			if err != nil {
				return err
			}
			c.logInitStage(contract)

			resp, err := c.SendContractPublishTx(cCtx, bytecode, contract.MaxGasLimit())
			if err != nil {
				return err
			}
			c.logPublishedStage(contract, resp)

			return nil
		},
	}
}

func (c *Container) erc20Demurrage() *cli.Command {
	return &cli.Command{
		Name:    "erc20-demurrage",
		Aliases: []string{"det"},
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
			&cli.Int64Flag{
				Name:     "demurrage-rate",
				Aliases:  []string{"expiry-rate"},
				Usage:    "This is the rate at which the voucher will expire per redistribution period",
				Value:    2,
				Required: false,
				Action: func(ctx *cli.Context, x int64) error {
					if x < 1 {
						return errors.New("rate should be atleast 1 percent")
					}
					return nil
				},
			},
			&cli.Int64Flag{
				Name:     "redistribution-period",
				Aliases:  []string{"period-minutes"},
				Usage:    "Number of minutes between each time the demurraged value can be withdrawn to the Sink Account",
				Value:    43200,
				Required: false,
				Action: func(ctx *cli.Context, x int64) error {
					if x < 10080 {
						return errors.New("redistribution period should be atleast 1 week equivalent in minutes")
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:     "sink-address",
				Aliases:  []string{"community-fund"},
				Usage:    "The initial Sink Address",
				Required: true,
			},
		},
		Action: func(cCtx *cli.Context) error {
			decayLevel, err := util.CalculateDecayLevel(
				cCtx.Int64("demurrage-rate"),
				cCtx.Int64("redistribution-period"),
			)
			if err != nil {
				return err
			}

			c.Logg.Info("calculated decay level", "decay_level", decayLevel.String())

			contract := contract.NewERC20Demurrage(contract.ERC20DemurrageConstructorArgs{
				Name:               cCtx.String("name"),
				Symbol:             strings.ToUpper(cCtx.String("symbol")),
				Decimals:           uint8(cCtx.Uint("decimals")),
				DecayLevel:         decayLevel,
				PeriodMinutes:      big.NewInt(int64(cCtx.Uint64("redistribution-period"))),
				DefaultSinkAddress: ethutils.HexToAddress(cCtx.String("sink-address")),
			})
			bytecode, err := contract.Bytecode()
			if err != nil {
				return err
			}
			c.logInitStage(contract)

			resp, err := c.SendContractPublishTx(cCtx, bytecode, contract.MaxGasLimit())
			if err != nil {
				return err
			}
			c.logPublishedStage(contract, resp)

			return nil
		},
	}
}

func (c *Container) periodSimple() *cli.Command {
	return &cli.Command{
		Name:    "period-simple",
		Aliases: []string{"period"},
		Usage:   "Publish the period backend smart contract to complement the gas faucet contract",
		Action: func(cCtx *cli.Context) error {
			contract := contract.NewPeriodSimple()
			bytecode, err := contract.Bytecode()
			if err != nil {
				return err
			}
			c.logInitStage(contract)

			resp, err := c.SendContractPublishTx(cCtx, bytecode, contract.MaxGasLimit())
			if err != nil {
				return err
			}
			c.logPublishedStage(contract, resp)

			return nil
		},
	}
}

func (c *Container) ethFaucet() *cli.Command {
	return &cli.Command{
		Name:    "eth-faucet",
		Aliases: []string{"faucet"},
		Usage:   "Publish the gas faucet smart contract",
		Action: func(cCtx *cli.Context) error {
			contract := contract.NewEthFaucet()
			bytecode, err := contract.Bytecode()
			if err != nil {
				return err
			}
			c.logInitStage(contract)

			resp, err := c.SendContractPublishTx(cCtx, bytecode, contract.MaxGasLimit())
			if err != nil {
				return err
			}
			c.logPublishedStage(contract, resp)

			return nil
		},
	}
}

func (c *Container) accountsIndex() *cli.Command {
	return &cli.Command{
		Name:    "accounts-index",
		Aliases: []string{"user-index"},
		Usage:   "Publish the accounts index smart contract",
		Action: func(cCtx *cli.Context) error {
			contract := contract.NewAccountsIndex()
			bytecode, err := contract.Bytecode()
			if err != nil {
				return err
			}
			c.logInitStage(contract)

			resp, err := c.SendContractPublishTx(cCtx, bytecode, contract.MaxGasLimit())
			if err != nil {
				return err
			}
			c.logPublishedStage(contract, resp)

			return nil
		},
	}
}

func (c *Container) erc20() *cli.Command {
	return &cli.Command{
		Name:    "erc20",
		Aliases: []string{"giftable"},
		Usage:   "Publish the ERC20 (non-demurrage) smart contract",
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
				Name:     "expiry-timestamp",
				Aliases:  []string{"expiry"},
				Usage:    "Date time after which the tokens won't be transferable",
				Value:    0,
				Required: false,
			},
		},
		Action: func(cCtx *cli.Context) error {
			contract := contract.NewERC20(contract.ERC20ConstructorArgs{
				Name:            cCtx.String("name"),
				Symbol:          strings.ToUpper(cCtx.String("symbol")),
				Decimals:        uint8(cCtx.Uint("decimals")),
				ExpiryTimestamp: big.NewInt(int64(cCtx.Uint64("expiry-timestamp"))),
			})
			bytecode, err := contract.Bytecode()
			if err != nil {
				return err
			}
			c.logInitStage(contract)

			resp, err := c.SendContractPublishTx(cCtx, bytecode, contract.MaxGasLimit())
			if err != nil {
				return err
			}
			c.logPublishedStage(contract, resp)

			return nil
		},
	}
}

func (c *Container) contractsRegistry() *cli.Command {
	return &cli.Command{
		Name:    "contracts-registry",
		Aliases: []string{"registry"},
		Usage:   "Publish the contracts registry smart contract",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:     "identifier",
				Usage:    "Contract identifier",
				Required: true,
			},
		},
		Action: func(cCtx *cli.Context) error {
			contract := contract.NewContractsRegistry(cCtx.StringSlice("identifier"))
			bytecode, err := contract.Bytecode()
			if err != nil {
				return err
			}
			c.logInitStage(contract)

			resp, err := c.SendContractPublishTx(cCtx, bytecode, contract.MaxGasLimit())
			if err != nil {
				return err
			}
			c.logPublishedStage(contract, resp)

			return nil
		},
	}
}

func (c *Container) custodialRegistrationProxy() *cli.Command {
	return &cli.Command{
		Name:    "custodial-registration-proxy",
		Aliases: []string{"custodial"},
		Usage:   "Publish the custodial registration proxy smart contract",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "eth-faucet-address",
				Aliases:  []string{"gas-faucet", "faucet"},
				Usage:    "The gas faucet address",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "accounts-index-address",
				Aliases:  []string{"user-index"},
				Usage:    "The accounts index address",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "system-account-address",
				Aliases:  []string{"system-account"},
				Usage:    "The system account address",
				Required: true,
			},
		},
		Action: func(cCtx *cli.Context) error {
			contract := contract.NewCustodialRegistrationProxy(contract.CustodialRegistrationProxyArgs{
				EthFaucetAddress:     ethutils.HexToAddress(cCtx.String("eth-faucet-address")),
				AccountsIndexAddress: ethutils.HexToAddress(cCtx.String("accounts-index-address")),
				SystemAccountAddress: ethutils.HexToAddress(cCtx.String("system-account-address")),
			})
			bytecode, err := contract.Bytecode()
			if err != nil {
				return err
			}
			c.logInitStage(contract)

			resp, err := c.SendContractPublishTx(cCtx, bytecode, contract.MaxGasLimit())
			if err != nil {
				return err
			}
			c.logPublishedStage(contract, resp)

			return nil
		},
	}
}

func (c *Container) logInitStage(contract contract.Contract) {
	c.Logg.Info(fmt.Sprintf("publishing %s contract", contract.Name()),
		"version", contract.Version(),
		"constructor_args", contract.ConstructorArgs(),
	)
	c.Logg.Debug(fmt.Sprintf("publishing %s contract", contract.Name()),
		"license", contract.License(),
		"source", contract.Source(),
		"solidity_version", contract.SolidityVersion(),
		"evm_fork", contract.EVMFork(),
	)
}

func (c *Container) logPublishedStage(contract contract.Contract, resp PublishTxResp) {
	c.Logg.Info(fmt.Sprintf("successfully published %s contract", contract.Name()),
		"contract_address", resp.ContractAddress.Hex(),
		"tx_hash", resp.TxHash.String(),
	)
}
