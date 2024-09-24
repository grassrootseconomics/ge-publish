package accountsindex

import (
	_ "embed"

	"github.com/grassrootseconomics/celoutils/v2"
	"github.com/grassrootseconomics/ge-publish/pkg/util"
)

type (
	AccountsIndex struct {
		Constructor []any
	}
)

const (
	name            = "AccountsIndex"
	version         = "v0.6.0"
	license         = "AGPL-3.0"
	source          = "https://git.grassecon.net/cicnet/eth-accounts-index/src/branch/dev-0.6.0/solidity"
	solidityVersion = "0.8.25+commit.b61c2a91"
	evmFork         = "istanbul"

	gasLimit = 2_000_000
)

var (
	//go:embed AccountsIndex.bin
	bin string
	//go:embed AccountsIndex.json
	abi string
)

func (c *AccountsIndex) Name() string {
	return name
}

func (c *AccountsIndex) Version() string {
	return version
}

func (c *AccountsIndex) License() string {
	return license
}

func (c *AccountsIndex) Source() string {
	return source
}

func (c *AccountsIndex) SolidityVersion() string {
	return solidityVersion
}

func (c *AccountsIndex) EVMFork() string {
	return evmFork
}

func (c *AccountsIndex) ConstructorArgs() []string {
	return util.DumpConstructorArgs(c.Constructor)
}

func (c *AccountsIndex) Bytecode() ([]byte, error) {
	return celoutils.PrepareContractBytecodeData(bin, abi, c.Constructor)
}

func (c *AccountsIndex) MaxGasLimit() uint64 {
	return gasLimit
}
