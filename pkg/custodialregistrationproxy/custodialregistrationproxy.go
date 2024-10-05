package custodialregistrationproxy

import (
	_ "embed"

	"github.com/grassrootseconomics/ethutils"
	"github.com/grassrootseconomics/ge-publish/pkg/util"
)

type (
	CustodialRegistrationProxy struct {
		Constructor []any
	}
)

const (
	name            = "CustodialRegistrationProxy"
	version         = "v1.0.0"
	license         = "AGPL-3.0"
	source          = "https://github.com/grassrootseconomics/custodial-registration-proxy/tree/master/solidity"
	solidityVersion = "0.8.25+commit.b61c2a91"
	evmFork         = "istanbul"

	gasLimit = 750_000
)

var (
	//go:embed CustodialRegistrationProxy.bin
	bin string
	//go:embed CustodialRegistrationProxy.json
	abi string
)

func (c *CustodialRegistrationProxy) Name() string {
	return name
}

func (c *CustodialRegistrationProxy) Version() string {
	return version
}

func (c *CustodialRegistrationProxy) License() string {
	return license
}

func (c *CustodialRegistrationProxy) Source() string {
	return source
}

func (c *CustodialRegistrationProxy) SolidityVersion() string {
	return solidityVersion
}

func (c *CustodialRegistrationProxy) EVMFork() string {
	return evmFork
}

func (c *CustodialRegistrationProxy) ConstructorArgs() []string {
	return util.DumpConstructorArgs(c.Constructor)
}

func (c *CustodialRegistrationProxy) Bytecode() ([]byte, error) {
	return ethutils.PrepareContractBytecodeData(bin, abi, c.Constructor)
}

func (c *CustodialRegistrationProxy) MaxGasLimit() uint64 {
	return gasLimit
}
