package limiterindex

import (
	_ "embed"

	"github.com/celo-org/celo-blockchain/common"
	"github.com/grassrootseconomics/celoutils/v2"
	"github.com/grassrootseconomics/ge-publish/pkg/contract"
)

type (
	LimiterIndex struct{}

	LimiterIndexConstructorArgs struct {
		Holder         common.Address
		LimiterAddress common.Address
	}
)

const (
	version = "0.0.1"

	gasLimit = 750_000
)

var (
	//go:embed LimiterIndex.bin
	bin string
	//go:embed LimiterIndex.json
	abi string
)

func NewLimiterIndexContract() contract.Contract[LimiterIndexConstructorArgs] {
	return &LimiterIndex{}
}

func (c *LimiterIndex) Version() string {
	return version
}

func (c *LimiterIndex) GasLimit() uint64 {
	return gasLimit
}

func (c *LimiterIndex) Bytecode(args LimiterIndexConstructorArgs) ([]byte, error) {
	return celoutils.PrepareContractBytecodeData(bin, abi, []any{
		args.Holder,
		args.LimiterAddress,
	})
}
