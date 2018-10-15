package money

type calculator struct{}

func (c *calculator) add(a, b *Amount) *Amount {
	return &Amount{a.Value() + b.Value()}
}

func (c *calculator) subtract(a, b *Amount) *Amount {
	return &Amount{a.Value() - b.Value()}
}

func (c *calculator) multiply(a *Amount, m int64) *Amount {
	return &Amount{a.Value() * m}
}

func (c *calculator) allocate(a *Amount, ratio, total int) *Amount {
	return &Amount{a.Value() * int64(ratio) / int64(total)}
}
