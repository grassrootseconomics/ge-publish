package swappool

import (
	_ "embed"

	"github.com/grassrootseconomics/ethutils"
	"github.com/grassrootseconomics/ge-publish/pkg/util"
)

type (
	SwapPool struct {
		Constructor []any
	}
)

const (
	name            = "SwapPool"
	version         = "5189aec"
	license         = "AGPL-3.0"
	source          = "https://github.com/grassrootseconomics/erc20-pool/blob/sohail/pool-updates/solidity/SwapPool.sol"
	solidityVersion = "0.8.25+commit.b61c2a91"
	evmFork         = "istanbul"

	gasLimit = 3_000_000
)

var (
	//go:embed SwapPool.bin
	bin string
	//go:embed SwapPool.json
	abi string
)

func (c *SwapPool) Name() string {
	return name
}

func (c *SwapPool) Version() string {
	return version
}

func (c *SwapPool) License() string {
	return license
}

func (c *SwapPool) Source() string {
	return source
}

func (c *SwapPool) SolidityVersion() string {
	return solidityVersion
}

func (c *SwapPool) EVMFork() string {
	return evmFork
}

func (c *SwapPool) ConstructorArgs() []string {
	return util.DumpConstructorArgs(c.Constructor)
}

func (c *SwapPool) Bytecode() ([]byte, error) {
	return ethutils.PrepareContractBytecodeData(bin, abi, c.Constructor)
}

func (c *SwapPool) MaxGasLimit() uint64 {
	return gasLimit
}
