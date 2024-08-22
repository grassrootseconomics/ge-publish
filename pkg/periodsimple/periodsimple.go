package periodsimple

import (
	_ "embed"

	"github.com/grassrootseconomics/celoutils/v2"
	"github.com/grassrootseconomics/ge-publish/pkg/util"
)

type (
	PeriodSimple struct {
		Constructor []any
	}
)

const (
	name            = "PeriodSimple"
	version         = "v0.4.1"
	license         = "AGPL-3.0"
	source          = "https://git.grassecon.net/cicnet/eth-faucet/src/branch/master/solidity"
	solidityVersion = "0.8.25+commit.b61c2a91"
	evmFork         = "istanbul"

	gasLimit = 1_000_000
)

var (
	//go:embed PeriodSimple.bin
	bin string
	//go:embed PeriodSimple.json
	abi string
)

func (c *PeriodSimple) Name() string {
	return name
}

func (c *PeriodSimple) Version() string {
	return version
}

func (c *PeriodSimple) License() string {
	return license
}

func (c *PeriodSimple) Source() string {
	return source
}

func (c *PeriodSimple) SolidityVersion() string {
	return solidityVersion
}

func (c *PeriodSimple) EVMFork() string {
	return evmFork
}

func (c *PeriodSimple) ConstructorArgs() []string {
	return util.DumpConstructorArgs(c.Constructor)
}

func (c *PeriodSimple) Bytecode() ([]byte, error) {
	return celoutils.PrepareContractBytecodeData(bin, abi, c.Constructor)
}

func (c *PeriodSimple) MaxGasLimit() uint64 {
	return gasLimit
}
