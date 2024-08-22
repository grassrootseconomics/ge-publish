package erc20demurrage

import (
	_ "embed"

	"github.com/grassrootseconomics/celoutils/v2"
	"github.com/grassrootseconomics/ge-publish/pkg/util"
)

type (
	ERC20Demurrage struct {
		Constructor []any
	}
)

const (
	name            = "ERC20Demurrage"
	version         = "v0.5.7"
	license         = "AGPL-3.0"
	source          = "https://github.com/nolash/erc20-demurrage-token/tree/dev-0.5.7/solidity"
	solidityVersion = "0.8.25+commit.b61c2a91"
	evmFork         = "istanbul"

	gasLimit = 6_000_000
)

var (
	//go:embed ERC20Demurrage.bin
	bin string
	//go:embed ERC20Demurrage.json
	abi string
)

func (c *ERC20Demurrage) Name() string {
	return name
}

func (c *ERC20Demurrage) Version() string {
	return version
}

func (c *ERC20Demurrage) License() string {
	return license
}

func (c *ERC20Demurrage) Source() string {
	return source
}

func (c *ERC20Demurrage) SolidityVersion() string {
	return solidityVersion
}

func (c *ERC20Demurrage) EVMFork() string {
	return evmFork
}

func (c *ERC20Demurrage) ConstructorArgs() []string {
	return util.DumpConstructorArgs(c.Constructor)
}

func (c *ERC20Demurrage) Bytecode() ([]byte, error) {
	return celoutils.PrepareContractBytecodeData(bin, abi, c.Constructor)
}

func (c *ERC20Demurrage) MaxGasLimit() uint64 {
	return gasLimit
}
