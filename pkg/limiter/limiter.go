package limiter

import (
	_ "embed"

	"github.com/grassrootseconomics/ethutils"
	"github.com/grassrootseconomics/ge-publish/pkg/util"
)

type (
	Limiter struct {
		Constructor []any
	}
)

const (
	name             = "Limiter"
	version          = "d5c38bf"
	license          = "AGPL-3.0"
	source           = "https://github.com/grassrootseconomics/erc20-limiter/blob/master/solidity/Limiter.sol"
	solidityVersion  = "0.8.30+commit.73712a01"
	evmFork          = "shanghai"
	optimizationRuns = 200
	optimized        = true

	gasLimit = 1_000_000
)

var (
	//go:embed Limiter.bin
	bin string
	//go:embed Limiter.json
	abi string
)

func (c *Limiter) Name() string {
	return name
}

func (c *Limiter) Version() string {
	return version
}

func (c *Limiter) License() string {
	return license
}

func (c *Limiter) Source() string {
	return source
}

func (c *Limiter) SolidityVersion() string {
	return solidityVersion
}

func (c *Limiter) EVMFork() string {
	return evmFork
}

func (c *Limiter) ConstructorArgs() []string {
	return util.DumpConstructorArgs(c.Constructor)
}

func (c *Limiter) Bytecode() ([]byte, error) {
	return ethutils.PrepareContractBytecodeData(bin, abi, c.Constructor)
}

func (c *Limiter) MaxGasLimit() uint64 {
	return gasLimit
}
