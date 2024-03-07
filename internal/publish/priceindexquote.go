package publish

import (
	"fmt"

	"github.com/grassrootseconomics/ge-publish/internal/provider"
	"github.com/grassrootseconomics/ge-publish/pkg/priceindexquote"
	"github.com/urfave/cli/v2"
)

func (c *Command) RegisterPriceIndexQuoterCommand() *cli.Command {
	return &cli.Command{
		Name:  "price-index-quoter",
		Usage: "Publish the price index quoter smart contract",
		Action: func(cCtx *cli.Context) error {
			contract := priceindexquote.NewPriceIndexQuoterContract()
			bytecode, err := contract.Bytecode(priceindexquote.PriceIndexQuoterConstructorArgs{})
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
