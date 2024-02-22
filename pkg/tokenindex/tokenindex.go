package tokenindex

import (
	_ "embed"

	"github.com/grassrootseconomics/celoutils/v2"
	"github.com/grassrootseconomics/ge-publish/pkg/contract"
)

type (
	TokenIndex struct{}

	TokenIndexConstructorArgs struct{}
)

const (
	version = "0.6.3"

	gasLimit = 2_000_000
)

var (
	//go:embed TokenIndex.bin
	bin string
	//go:embed TokenIndex.json
	abi string
)

func NewTokenIndexContract() contract.Contract[TokenIndexConstructorArgs] {
	return &TokenIndex{}
}

func (c *TokenIndex) Version() string {
	return version
}

func (c *TokenIndex) GasLimit() uint64 {
	return gasLimit
}

func (c *TokenIndex) Bytecode(args TokenIndexConstructorArgs) ([]byte, error) {
	return celoutils.PrepareContractBytecodeData(bin, abi, []any{})
}
