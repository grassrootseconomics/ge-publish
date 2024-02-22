package publish

import (
	"fmt"

	"github.com/grassrootseconomics/ge-publish/internal/provider"
	"github.com/grassrootseconomics/ge-publish/pkg/limiter"
	"github.com/urfave/cli/v2"
)

func (c *Command) RegisterLimiterCommand() *cli.Command {
	return &cli.Command{
		Name:  "limiter",
		Usage: "Publish the limiter smart contract",
		Action: func(cCtx *cli.Context) error {
			contract := limiter.NewLimiterContract()
			bytecode, err := contract.Bytecode(limiter.LimiterConstructorArgs{})
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
