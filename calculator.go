package money

// Calculator calculates different amounts
type Calculator struct{}

// Add adds two amounts
func (c *Calculator) Add(a, b *Amount) *Amount {
	return &Amount{a.val + b.val}
}

// Subtract subtracts one amount from an other
func (c *Calculator) Subtract(a, b *Amount) *Amount {
	return &Amount{a.val - b.val}
}
