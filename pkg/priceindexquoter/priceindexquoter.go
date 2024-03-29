package priceindexquoter

import (
	_ "embed"

	"github.com/grassrootseconomics/celoutils/v2"
	"github.com/grassrootseconomics/ge-publish/pkg/contract"
)

type (
	PriceIndexQuoter struct{}

	PriceIndexQuoterConstructorArgs struct{}
)

const (
	version = "0.1.0"

	gasLimit = 2_000_000
)

var (
	//go:embed PriceIndexQuoter.bin
	bin string
	//go:embed PriceIndexQuoter.json
	abi string
)

func NewPriceIndexQuoterContract() contract.Contract[PriceIndexQuoterConstructorArgs] {
	return &PriceIndexQuoter{}
}

func (c *PriceIndexQuoter) Version() string {
	return version
}

func (c *PriceIndexQuoter) GasLimit() uint64 {
	return gasLimit
}

func (c *PriceIndexQuoter) Bytecode(args PriceIndexQuoterConstructorArgs) ([]byte, error) {
	return celoutils.PrepareContractBytecodeData(bin, abi, []any{})
}
