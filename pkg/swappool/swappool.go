package swappool

import (
	_ "embed"

	"github.com/celo-org/celo-blockchain/common"
	"github.com/grassrootseconomics/celoutils/v2"
	"github.com/grassrootseconomics/ge-publish/pkg/contract"
)

type (
	SwapPool struct{}

	SwapPoolConstructorArgs struct {
		Name          string
		Symbol        string
		Decimals      uint8
		TokenRegistry common.Address
		TokenLimiter  common.Address
	}
)

const version = "0.0.1"

var (
	//go:embed SwapPool.bin
	bin string
	//go:embed SwapPool.json
	abi string
)

func NewSwapPoolContract() contract.Contract[SwapPoolConstructorArgs] {
	return &SwapPool{}
}

func (c *SwapPool) Version() string {
	return version
}

func (c *SwapPool) Bytecode(args SwapPoolConstructorArgs) ([]byte, error) {
	return celoutils.PrepareContractBytecodeData(bin, abi, []any{
		args.Name,
		args.Symbol,
		args.Decimals,
		args.TokenRegistry,
		args.TokenLimiter,
	})
}
