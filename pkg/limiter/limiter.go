package limiter

import (
	_ "embed"

	"github.com/grassrootseconomics/celoutils/v2"
	"github.com/grassrootseconomics/ge-publish/pkg/util"
)

type (
	Limiter struct {
		Constructor []any
	}
)

const (
	name            = "Limiter"
	version         = "v0.0.1"
	license         = "AGPL-3.0"
	source          = "https://github.com/nolash/erc20-limiter/blob/master/solidity/Limiter.sol"
	solidityVersion = "0.8.19+commit.7dd6d404"
	evmFork         = "byzantium"

	gasLimit = 750_000
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
	return celoutils.PrepareContractBytecodeData(bin, abi, c.Constructor)
}

func (c *Limiter) MaxGasLimit() uint64 {
	return gasLimit
}
