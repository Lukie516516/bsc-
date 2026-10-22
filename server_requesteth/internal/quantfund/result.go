package quantfund

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

var resultRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func CalcProfit(quantAmount string, profitBps int) (profitAmount string, returnAmount string, err error) {
	quantWei, err := DecimalToWei(quantAmount, 18)
	if err != nil {
		return "", "", err
	}
	profitWei := new(big.Int).Mul(quantWei, big.NewInt(int64(profitBps)))
	profitWei.Div(profitWei, big.NewInt(10000))
	returnWei := new(big.Int).Add(new(big.Int).Set(quantWei), profitWei)
	return WeiToDecimal(profitWei, 18), WeiToDecimal(returnWei, 18), nil
}

func NormalizeResultMode(mode string) string {
	if mode == "" {
		return "simulate"
	}
	return mode
}

func BuildDryRunTxHash(prefix, taskNo string) string {
	return fmt.Sprintf("%s-%s", prefix, taskNo)
}

func PickProfitBps(mode string, fixedBps, minBps, maxBps int) int {
	switch NormalizeResultMode(mode) {
	case "fixed":
		return fixedBps
	case "random_range":
		if minBps == maxBps {
			return minBps
		}
		if minBps > maxBps {
			minBps, maxBps = maxBps, minBps
		}
		return minBps + resultRand.Intn(maxBps-minBps+1)
	default:
		return fixedBps
	}
}
