package publish

import (
	"fmt"

	"github.com/grassrootseconomics/ge-publish/internal/provider"
	"github.com/grassrootseconomics/ge-publish/pkg/decimalquote"
	"github.com/urfave/cli/v2"
)

func (c *Command) RegisterDecimalQuoteCommand() *cli.Command {
	return &cli.Command{
		Name:  "decimal-quote",
		Usage: "Publish the decimal quote smart contract",
		Action: func(cCtx *cli.Context) error {
			contract := decimalquote.NewDecimalQuoteContract()
			bytecode, err := contract.Bytecode(decimalquote.DecimalQuoteConstructorArgs{})
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
