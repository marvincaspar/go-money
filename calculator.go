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

func (c *calculator) divide(a *Amount, d int64) *Amount {
	return &Amount{a.Value() / d}
}

func (c *calculator) modulus(a *Amount, d int64) *Amount {
	return &Amount{a.Value() % d}
}

func (c *calculator) allocate(a *Amount, ratio, total int) *Amount {
	return &Amount{a.Value() * int64(ratio) / int64(total)}
}

func (c *calculator) absolute(a *Amount) *Amount {
	if a.Value() < 0 {
		return &Amount{-a.Value()}
	}

	return &Amount{a.Value()}
}

func (c *calculator) negative(a *Amount) *Amount {
	if a.Value() > 0 {
		return &Amount{-a.Value()}
	}

	return &Amount{a.Value()}
}
