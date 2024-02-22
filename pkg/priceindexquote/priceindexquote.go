package priceindexquote

import (
	_ "embed"

	"github.com/grassrootseconomics/celoutils/v2"
	"github.com/grassrootseconomics/ge-publish/pkg/contract"
)

type (
	PriceIndexQuote struct{}

	PriceIndexQuoteConstructorArgs struct{}
)

const (
	version = "0.0.1"

	gasLimit = 1_000_000
)

var (
	//go:embed PriceIndexQuote.bin
	bin string
	//go:embed PriceIndexQuote.json
	abi string
)

func NewPriceIndexQuoteContract() contract.Contract[PriceIndexQuoteConstructorArgs] {
	return &PriceIndexQuote{}
}

func (c *PriceIndexQuote) Version() string {
	return version
}

func (c *PriceIndexQuote) GasLimit() uint64 {
	return gasLimit
}

func (c *PriceIndexQuote) Bytecode(args PriceIndexQuoteConstructorArgs) ([]byte, error) {
	return celoutils.PrepareContractBytecodeData(bin, abi, []any{})
}
