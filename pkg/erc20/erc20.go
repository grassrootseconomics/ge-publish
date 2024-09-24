package erc20

import (
	_ "embed"

	"github.com/grassrootseconomics/ethutils"
	"github.com/grassrootseconomics/ge-publish/pkg/util"
)

type (
	ERC20 struct {
		Constructor []any
	}
)

const (
	name            = "ERC20"
	version         = "v0.9.0"
	license         = "AGPL-3.0"
	source          = "https://git.grassecon.net/cicnet/eth-erc20/src/branch/master/solidity"
	solidityVersion = "0.8.25+commit.b61c2a91"
	evmFork         = "istanbul"

	gasLimit = 3_000_000
)

var (
	//go:embed ERC20.bin
	bin string
	//go:embed ERC20.json
	abi string
)

func (c *ERC20) Name() string {
	return name
}

func (c *ERC20) Version() string {
	return version
}

func (c *ERC20) License() string {
	return license
}

func (c *ERC20) Source() string {
	return source
}

func (c *ERC20) SolidityVersion() string {
	return solidityVersion
}

func (c *ERC20) EVMFork() string {
	return evmFork
}

func (c *ERC20) ConstructorArgs() []string {
	return util.DumpConstructorArgs(c.Constructor)
}

func (c *ERC20) Bytecode() ([]byte, error) {
	return ethutils.PrepareContractBytecodeData(bin, abi, c.Constructor)
}

func (c *ERC20) MaxGasLimit() uint64 {
	return gasLimit
}
