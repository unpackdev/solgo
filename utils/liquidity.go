package utils

type LiquidityProvider string

func (lp LiquidityProvider) String() string {
	return string(lp)
}

const (
	PinksaleLiquidtyProvider LiquidityProvider = "pinksale"
)
