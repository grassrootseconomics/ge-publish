package contractsregistry

import (
	_ "embed"

	"github.com/grassrootseconomics/ethutils"
	"github.com/grassrootseconomics/ge-publish/pkg/util"
)

type (
	ContractsRegistry struct {
		Constructor []any
	}
)

const (
	name            = "ContractsRegistry"
	version         = "v0.11.0"
	license         = "AGPL-3.0"
	source          = "https://git.grassecon.net/cicnet/eth-contract-registry/src/branch/dev-0.11.0/solidity"
	solidityVersion = "0.8.25+commit.b61c2a91"
	evmFork         = "istanbul"

	gasLimit = 2_000_000
)

var (
	//go:embed ContractsRegistry.bin
	bin string
	//go:embed ContractsRegistry.json
	abi string
)

func (c *ContractsRegistry) Name() string {
	return name
}

func (c *ContractsRegistry) Version() string {
	return version
}

func (c *ContractsRegistry) License() string {
	return license
}

func (c *ContractsRegistry) Source() string {
	return source
}

func (c *ContractsRegistry) SolidityVersion() string {
	return solidityVersion
}

func (c *ContractsRegistry) EVMFork() string {
	return evmFork
}

func (c *ContractsRegistry) ConstructorArgs() []string {
	return util.DumpConstructorArgs(c.Constructor)
}

func (c *ContractsRegistry) Bytecode() ([]byte, error) {
	return ethutils.PrepareContractBytecodeData(bin, abi, c.Constructor)
}

func (c *ContractsRegistry) MaxGasLimit() uint64 {
	return gasLimit
}
