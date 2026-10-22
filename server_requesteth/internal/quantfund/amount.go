package quantfund

import (
	"fmt"
	"math/big"
	"strings"
)

var ten = big.NewInt(10)

func DecimalToWei(value string, decimals int) (*big.Int, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return big.NewInt(0), nil
	}

	negative := false
	if strings.HasPrefix(value, "-") {
		negative = true
		value = strings.TrimPrefix(value, "-")
	}

	parts := strings.SplitN(value, ".", 2)
	intPart := parts[0]
	fracPart := ""
	if len(parts) == 2 {
		fracPart = parts[1]
	}
	if intPart == "" {
		intPart = "0"
	}
	if len(fracPart) > decimals {
		fracPart = fracPart[:decimals]
	}
	for len(fracPart) < decimals {
		fracPart += "0"
	}

	full := strings.TrimLeft(intPart+fracPart, "0")
	if full == "" {
		return big.NewInt(0), nil
	}

	out, ok := new(big.Int).SetString(full, 10)
	if !ok {
		return nil, fmt.Errorf("invalid decimal value: %s", value)
	}
	if negative {
		out.Neg(out)
	}
	return out, nil
}

func WeiToDecimal(value *big.Int, decimals int) string {
	if value == nil || value.Sign() == 0 {
		return "0"
	}
	if decimals <= 0 {
		return value.String()
	}

	negative := value.Sign() < 0
	if negative {
		value = new(big.Int).Abs(value)
	}

	base := new(big.Int).Exp(ten, big.NewInt(int64(decimals)), nil)
	intPart := new(big.Int).Quo(value, base)
	fracPart := new(big.Int).Mod(value, base)

	frac := fracPart.String()
	for len(frac) < decimals {
		frac = "0" + frac
	}
	frac = strings.TrimRight(frac, "0")

	res := intPart.String()
	if frac != "" {
		res += "." + frac
	}
	if negative {
		res = "-" + res
	}
	return res
}

func SplitByBps(total *big.Int, bps int) (*big.Int, *big.Int) {
	if total == nil || total.Sign() == 0 {
		return big.NewInt(0), big.NewInt(0)
	}
	quant := new(big.Int).Mul(total, big.NewInt(int64(bps)))
	quant.Div(quant, big.NewInt(10000))
	reserve := new(big.Int).Sub(new(big.Int).Set(total), quant)
	return quant, reserve
}
