package billing_tools

import "github.com/shopspring/decimal"

func Add(a, b decimal.NullDecimal) decimal.Decimal {
	if a.Valid && b.Valid {
		return a.Decimal.Add(b.Decimal)
	}

	if a.Valid {
		return a.Decimal
	}

	if b.Valid {
		return b.Decimal
	}

	return decimal.NewFromInt(0)
}
