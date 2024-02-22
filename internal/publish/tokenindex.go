package publish

import (
	"fmt"

	"github.com/grassrootseconomics/ge-publish/internal/provider"
	"github.com/grassrootseconomics/ge-publish/pkg/tokenindex"
	"github.com/urfave/cli/v2"
)

func (c *Command) RegisterTokenIndexCommand() *cli.Command {
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

			resp, err := provider.SendContractPublishTx(cCtx, bytecode, contract.GasLimit())
			if err != nil {
				return err
			}

			fmt.Println(resp.TxHash.String())

			return nil
		},
	}
}
