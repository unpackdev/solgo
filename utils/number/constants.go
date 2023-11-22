package number

type Rounding int

const (
	RoundDown Rounding = iota
	RoundHalfUp
	RoundUp
)

// Valid check this rounding mode is valid
func (r Rounding) Valid() bool {
	return r == RoundDown ||
		r == RoundHalfUp ||
		r == RoundUp
}
