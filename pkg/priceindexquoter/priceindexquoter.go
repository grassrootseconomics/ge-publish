package priceindexquoter

import (
	_ "embed"

	"github.com/grassrootseconomics/ethutils"
	"github.com/grassrootseconomics/ge-publish/pkg/util"
)

type (
	PriceIndexQuoter struct {
		Constructor []any
	}
)

const (
	name             = "PriceIndexQuoter"
	version          = "6721100"
	license          = "AGPL-3.0"
	source           = "https://github.com/grassrootseconomics/price-index-quoter/blob/master/src/PriceIndexQuoter.sol"
	solidityVersion  = "0.8.30+commit.73712a01"
	evmFork          = "shanghai"
	optimizationRuns = 200
	optimized        = true

	gasLimit = 1_000_000
)

var (
	//go:embed PriceIndexQuoter.bin
	bin string
	//go:embed PriceIndexQuoter.json
	abi string
)

func (c *PriceIndexQuoter) Name() string {
	return name
}

func (c *PriceIndexQuoter) Version() string {
	return version
}

func (c *PriceIndexQuoter) License() string {
	return license
}

func (c *PriceIndexQuoter) Source() string {
	return source
}

func (c *PriceIndexQuoter) SolidityVersion() string {
	return solidityVersion
}

func (c *PriceIndexQuoter) EVMFork() string {
	return evmFork
}

func (c *PriceIndexQuoter) ConstructorArgs() []string {
	return util.DumpConstructorArgs(c.Constructor)
}

func (c *PriceIndexQuoter) Bytecode() ([]byte, error) {
	return ethutils.PrepareContractBytecodeData(bin, abi, c.Constructor)
}

func (c *PriceIndexQuoter) MaxGasLimit() uint64 {
	return gasLimit
}
