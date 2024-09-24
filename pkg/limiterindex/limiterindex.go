package limiterindex

import (
	_ "embed"

	"github.com/ethereum/go-ethereum/common"
	"github.com/grassrootseconomics/ethutils"
	"github.com/grassrootseconomics/ge-publish/pkg/util"
)

type (
	LimiterIndex struct {
		Constructor []any
	}

	LimiterIndexConstructorArgs struct {
		Holder              common.Address
		LimiterIndexAddress common.Address
	}
)

const (
	name            = "LimiterIndex"
	version         = "v0.0.1"
	license         = "AGPL-3.0"
	source          = "https://github.com/nolash/erc20-limiter/blob/master/solidity/LimiterIndex.sol"
	solidityVersion = "0.8.19+commit.7dd6d404"
	evmFork         = "byzantium"

	gasLimit = 750_000
)

var (
	//go:embed LimiterIndex.bin
	bin string
	//go:embed LimiterIndex.json
	abi string
)

func (c *LimiterIndex) Name() string {
	return name
}

func (c *LimiterIndex) Version() string {
	return version
}

func (c *LimiterIndex) License() string {
	return license
}

func (c *LimiterIndex) Source() string {
	return source
}

func (c *LimiterIndex) SolidityVersion() string {
	return solidityVersion
}

func (c *LimiterIndex) EVMFork() string {
	return evmFork
}

func (c *LimiterIndex) ConstructorArgs() []string {
	return util.DumpConstructorArgs(c.Constructor)
}

func (c *LimiterIndex) Bytecode() ([]byte, error) {
	return ethutils.PrepareContractBytecodeData(bin, abi, c.Constructor)
}

func (c *LimiterIndex) MaxGasLimit() uint64 {
	return gasLimit
}
