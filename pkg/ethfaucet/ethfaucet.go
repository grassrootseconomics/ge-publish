package ethfaucet

import (
	_ "embed"

	"github.com/grassrootseconomics/ethutils"
	"github.com/grassrootseconomics/ge-publish/pkg/util"
)

type (
	EthFaucet struct {
		Constructor []any
	}
)

const (
	name            = "EthFaucet"
	version         = "v0.4.1"
	license         = "AGPL-3.0"
	source          = "https://git.grassecon.net/cicnet/eth-faucet/src/branch/master/solidity"
	solidityVersion = "0.8.25+commit.b61c2a91"
	evmFork         = "istanbul"

	gasLimit = 3_000_000
)

var (
	//go:embed EthFaucet.bin
	bin string
	//go:embed EthFaucet.json
	abi string
)

func (c *EthFaucet) Name() string {
	return name
}

func (c *EthFaucet) Version() string {
	return version
}

func (c *EthFaucet) License() string {
	return license
}

func (c *EthFaucet) Source() string {
	return source
}

func (c *EthFaucet) SolidityVersion() string {
	return solidityVersion
}

func (c *EthFaucet) EVMFork() string {
	return evmFork
}

func (c *EthFaucet) ConstructorArgs() []string {
	return util.DumpConstructorArgs(c.Constructor)
}

func (c *EthFaucet) Bytecode() ([]byte, error) {
	return ethutils.PrepareContractBytecodeData(bin, abi, c.Constructor)
}

func (c *EthFaucet) MaxGasLimit() uint64 {
	return gasLimit
}
