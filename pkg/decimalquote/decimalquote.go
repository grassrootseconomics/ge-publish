package decimalquote

import (
	_ "embed"

	"github.com/grassrootseconomics/ethutils"
	"github.com/grassrootseconomics/ge-publish/pkg/util"
)

type (
	DecimalQuote struct {
		Constructor []any
	}
)

const (
	name            = "DecimalQuote"
	version         = "7bc3212"
	license         = "AGPL-3.0"
	source          = "https://github.com/nolash/erc20-pool/blob/master/solidity/DecimalQuote.sol"
	solidityVersion = "0.8.19+commit.7dd6d404"
	evmFork         = "byzantium"

	gasLimit = 750_000
)

var (
	//go:embed DecimalQuote.bin
	bin string
	//go:embed DecimalQuote.json
	abi string
)

func (c *DecimalQuote) Name() string {
	return name
}

func (c *DecimalQuote) Version() string {
	return version
}

func (c *DecimalQuote) License() string {
	return license
}

func (c *DecimalQuote) Source() string {
	return source
}

func (c *DecimalQuote) SolidityVersion() string {
	return solidityVersion
}

func (c *DecimalQuote) EVMFork() string {
	return evmFork
}

func (c *DecimalQuote) ConstructorArgs() []string {
	return util.DumpConstructorArgs(c.Constructor)
}

func (c *DecimalQuote) Bytecode() ([]byte, error) {
	return ethutils.PrepareContractBytecodeData(bin, abi, c.Constructor)
}

func (c *DecimalQuote) MaxGasLimit() uint64 {
	return gasLimit
}
