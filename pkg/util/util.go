package util

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func DumpConstructorArgs(constructorArgs []any) []string {
	constructorArgsLen := len(constructorArgs)
	if constructorArgsLen < 1 {
		return []string{}
	}

	dumpedArgs := make([]string, len(constructorArgs))
	for i, arg := range constructorArgs {
		switch v := arg.(type) {
		case string:
			dumpedArgs[i] = v
		case uint8:
			dumpedArgs[i] = fmt.Sprintf("%d", v)
		case *big.Int:
			dumpedArgs[i] = v.String()
		case common.Address:
			dumpedArgs[i] = v.String()
		}
	}
	return dumpedArgs
}
