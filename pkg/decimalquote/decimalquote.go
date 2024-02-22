package decimalquote

import (
	_ "embed"

	"github.com/grassrootseconomics/celoutils/v2"
	"github.com/grassrootseconomics/ge-publish/pkg/contract"
)

type (
	DecimalQuote struct{}

	DecimalQuoteConstructorArgs struct{}
)

const (
	version = "0.0.1"

	gasLimit = 750_000
)

var (
	//go:embed DecimalQuote.bin
	bin string
	//go:embed DecimalQuote.json
	abi string
)

func NewDecimalQuoteContract() contract.Contract[DecimalQuoteConstructorArgs] {
	return &DecimalQuote{}
}

func (c *DecimalQuote) Version() string {
	return version
}

func (c *DecimalQuote) GasLimit() uint64 {
	return gasLimit
}

func (c *DecimalQuote) Bytecode(args DecimalQuoteConstructorArgs) ([]byte, error) {
	return celoutils.PrepareContractBytecodeData(bin, abi, []any{})
}
