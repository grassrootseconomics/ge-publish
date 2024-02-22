package limiter

import (
	_ "embed"

	"github.com/grassrootseconomics/celoutils/v2"
	"github.com/grassrootseconomics/ge-publish/pkg/contract"
)

type (
	Limiter struct{}

	LimiterConstructorArgs struct{}
)

const (
	version = "0.0.1"

	gasLimit = 750_000
)

var (
	//go:embed Limiter.bin
	bin string
	//go:embed Limiter.json
	abi string
)

func NewLimiterContract() contract.Contract[LimiterConstructorArgs] {
	return &Limiter{}
}

func (c *Limiter) Version() string {
	return version
}

func (c *Limiter) GasLimit() uint64 {
	return gasLimit
}

func (c *Limiter) Bytecode(args LimiterConstructorArgs) ([]byte, error) {
	return celoutils.PrepareContractBytecodeData(bin, abi, []any{})
}
