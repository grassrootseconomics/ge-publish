package contract

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/grassrootseconomics/ge-publish/pkg/accountsindex"
	"github.com/grassrootseconomics/ge-publish/pkg/contractsregistry"
	"github.com/grassrootseconomics/ge-publish/pkg/custodialregistrationproxy"
	"github.com/grassrootseconomics/ge-publish/pkg/decimalquote"
	"github.com/grassrootseconomics/ge-publish/pkg/erc20"
	"github.com/grassrootseconomics/ge-publish/pkg/erc20demurrage"
	"github.com/grassrootseconomics/ge-publish/pkg/ethfaucet"
	"github.com/grassrootseconomics/ge-publish/pkg/limiter"
	"github.com/grassrootseconomics/ge-publish/pkg/limiterindex"
	"github.com/grassrootseconomics/ge-publish/pkg/periodsimple"
	"github.com/grassrootseconomics/ge-publish/pkg/priceindexquoter"
	"github.com/grassrootseconomics/ge-publish/pkg/swappool"
	"github.com/grassrootseconomics/ge-publish/pkg/tokenindex"
)

type (
	Contract interface {
		Name() string
		// Version return the smart contract source version, either a tag or a short commit hash
		Version() string
		// License return an SPDX identifier of the license
		License() string
		// Source returns a URL pointing to the source file or directory
		Source() string
		// SolidityVersion returns the Solidity version with which the bytecode was compiled with
		SolidityVersion() string
		// EVM returns the evm fork flag passed to the solidity compiler during compilation
		EVMFork() string
		// ConstructorArgs return the passed constructor args
		ConstructorArgs() []string
		// Bytecode returns the compiled bytecode including the passed constructor args
		Bytecode() ([]byte, error)
		// MaxGasLimit returns the expected max gas the constract deployment is expected to use
		MaxGasLimit() uint64
	}

	SwapPoolConstructorArgs struct {
		Name          string
		Symbol        string
		Decimals      uint8
		TokenRegistry common.Address
		TokenLimiter  common.Address
	}

	LimiterIndexConstructorArgs struct {
		Holder         common.Address
		LimiterAddress common.Address
	}

	ERC20DemurrageConstructorArgs struct {
		Name               string
		Symbol             string
		Decimals           uint8
		DecayLevel         *big.Int
		PeriodMinutes      *big.Int
		DefaultSinkAddress common.Address
	}

	ERC20ConstructorArgs struct {
		Name            string
		Symbol          string
		Decimals        uint8
		ExpiryTimestamp *big.Int
	}

	CustodialRegistrationProxyArgs struct {
		EthFaucetAddress     common.Address
		AccountsIndexAddress common.Address
		SystemAccountAddress common.Address
	}
)

func NewDecimalQuote() Contract {
	return &decimalquote.DecimalQuote{
		Constructor: []any{},
	}
}

func NewSwapPool(args SwapPoolConstructorArgs) Contract {
	return &swappool.SwapPool{
		Constructor: []any{
			args.Name,
			args.Symbol,
			args.Decimals,
			args.TokenRegistry,
			args.TokenLimiter,
		},
	}
}

func NewLimiter() Contract {
	return &limiter.Limiter{
		Constructor: []any{},
	}
}

func NewLimiterIndex(args LimiterIndexConstructorArgs) Contract {
	return &limiterindex.LimiterIndex{
		Constructor: []any{
			args.Holder,
			args.LimiterAddress,
		},
	}
}

func NewPriceIndexQuoter() Contract {
	return &priceindexquoter.PriceIndexQuoter{
		Constructor: []any{},
	}
}

func NewTokenIndex() Contract {
	return &tokenindex.TokenIndex{
		Constructor: []any{},
	}
}

func NewERC20Demurrage(args ERC20DemurrageConstructorArgs) Contract {
	return &erc20demurrage.ERC20Demurrage{
		Constructor: []any{
			args.Name,
			args.Symbol,
			args.Decimals,
			args.DecayLevel,
			args.PeriodMinutes,
			args.DefaultSinkAddress,
		},
	}
}

func NewPeriodSimple() Contract {
	return &periodsimple.PeriodSimple{
		Constructor: []any{},
	}
}

func NewEthFaucet() Contract {
	return &ethfaucet.EthFaucet{
		Constructor: []any{},
	}
}

func NewAccountsIndex() Contract {
	return &accountsindex.AccountsIndex{
		Constructor: []any{},
	}
}

func NewERC20(args ERC20ConstructorArgs) Contract {
	return &erc20.ERC20{
		Constructor: []any{
			args.Name,
			args.Symbol,
			args.Decimals,
			args.ExpiryTimestamp,
		},
	}
}

func NewContractsRegistry(identifiers []string) Contract {
	// bytes32[] memory _identifiers
	byte32identifiers := []common.Hash{}
	for _, v := range identifiers {
		byte32identifiers = append(byte32identifiers, common.BytesToHash(common.RightPadBytes([]byte(v), 32)))
	}

	return &contractsregistry.ContractsRegistry{
		Constructor: []any{
			byte32identifiers,
		},
	}
}

func NewCustodialRegistrationProxy(args CustodialRegistrationProxyArgs) Contract {
	return &custodialregistrationproxy.CustodialRegistrationProxy{
		Constructor: []any{
			args.EthFaucetAddress,
			args.AccountsIndexAddress,
			args.SystemAccountAddress,
		},
	}
}
