package erc20demurrage

import (
	_ "embed"
	"math/big"

	"github.com/celo-org/celo-blockchain/common"
	"github.com/grassrootseconomics/celoutils/v2"
	"github.com/grassrootseconomics/ge-publish/pkg/contract"
)

type (
	ERC20Demurrage struct{}

	ERC20DemurrageConstructorArgs struct {
		Name               string
		Symbol             string
		Decimals           uint8
		DecayLevel         *big.Int
		PeriodMinutes      *big.Int
		DefaultSinkAddress common.Address
	}
)

const (
	version = "0.5.6"

	gasLimit = 6_000_000
)

var (
	//go:embed ERC20Demurrage.bin
	bin string
	//go:embed ERC20Demurrage.json
	abi string
)

func NewERC20DemurrageContract() contract.Contract[ERC20DemurrageConstructorArgs] {
	return &ERC20Demurrage{}
}

func (c *ERC20Demurrage) Version() string {
	return version
}

func (c *ERC20Demurrage) GasLimit() uint64 {
	return gasLimit
}

func (c *ERC20Demurrage) Bytecode(args ERC20DemurrageConstructorArgs) ([]byte, error) {
	return celoutils.PrepareContractBytecodeData(bin, abi, []any{
		args.Name,
		args.Symbol,
		args.Decimals,
		args.DecayLevel,
		args.PeriodMinutes,
		args.DefaultSinkAddress,
	})
}
