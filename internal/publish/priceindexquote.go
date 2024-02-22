package publish

import (
	"fmt"

	"github.com/grassrootseconomics/ge-publish/internal/provider"
	"github.com/grassrootseconomics/ge-publish/pkg/priceindexquote"
	"github.com/urfave/cli/v2"
)

func (c *Command) RegisterPriceIndexQuoteCommand() *cli.Command {
	return &cli.Command{
		Name:  "price-index-quote",
		Usage: "Publish the price index quote smart contract",
		Action: func(cCtx *cli.Context) error {
			contract := priceindexquote.NewPriceIndexQuoteContract()
			bytecode, err := contract.Bytecode(priceindexquote.PriceIndexQuoteConstructorArgs{})
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
