package money

// Amount stores the amounts used for calculations
type Amount struct {
	val int64
}

// Value returns the amount value
func (a *Amount) Value() int64 {
	return a.val
}
