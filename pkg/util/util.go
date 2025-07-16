package util

import (
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func CalculateDecayLevel(demurrageRate, redistributionPeriod int64) (*big.Int, error) {
	if redistributionPeriod <= 0 {
		return nil, errors.New("redistribution period must be positive")
	}

	// (1 - rate/100) ^ (1 / period)
	decayLevel := math.Pow(1.0-float64(demurrageRate)/100.0, 1.0/float64(redistributionPeriod))

	if decayLevel >= 1.0 {
		return nil, errors.New("demurrage level must be less than 100%")
	}

	// Convert to 64-bit fixed-point
	scaled := decayLevel * (1 << 64)
	return new(big.Int).SetUint64(uint64(scaled)), nil
}

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
		default:
			dumpedArgs[i] = fmt.Sprintf("%v", v)
		}
	}
	return dumpedArgs
}
