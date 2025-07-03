package tokenindex

import (
	_ "embed"

	"github.com/grassrootseconomics/ethutils"
	"github.com/grassrootseconomics/ge-publish/pkg/util"
)

type (
	TokenIndex struct {
		Constructor []any
	}
)

const (
	name             = "TokenIndex"
	version          = "v0.6.3"
	license          = "AGPL-3.0"
	source           = "https://github.com/nolash/eth-token-index/blob/master/solidity/TokenUniqueSymbolIndex.sol"
	solidityVersion  = "0.8.30+commit.73712a01"
	evmFork          = "shanghai"
	optimizationRuns = 200
	optimized        = true

	gasLimit = 1_500_000
)

var (
	//go:embed TokenIndex.bin
	bin string
	//go:embed TokenIndex.json
	abi string
)

func (c *TokenIndex) Name() string {
	return name
}

func (c *TokenIndex) Version() string {
	return version
}

func (c *TokenIndex) License() string {
	return license
}

func (c *TokenIndex) Source() string {
	return source
}

func (c *TokenIndex) SolidityVersion() string {
	return solidityVersion
}

func (c *TokenIndex) EVMFork() string {
	return evmFork
}

func (c *TokenIndex) ConstructorArgs() []string {
	return util.DumpConstructorArgs(c.Constructor)
}

func (c *TokenIndex) Bytecode() ([]byte, error) {
	return ethutils.PrepareContractBytecodeData(bin, abi, c.Constructor)
}

func (c *TokenIndex) MaxGasLimit() uint64 {
	return gasLimit
}
