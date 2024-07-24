package billing_tools

import (
	"github.com/shopspring/decimal"
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    decimal.NullDecimal
		b    decimal.NullDecimal
		want decimal.Decimal
	}{
		{
			name: "both valid",
			a:    decimal.NullDecimal{Decimal: decimal.NewFromInt(1), Valid: true},
			b:    decimal.NullDecimal{Decimal: decimal.NewFromInt(2), Valid: true},
			want: decimal.NewFromInt(3),
		},
		{
			name: "a valid",
			a:    decimal.NullDecimal{Decimal: decimal.NewFromInt(1), Valid: true},
			b:    decimal.NullDecimal{Decimal: decimal.NewFromInt(0), Valid: false},
			want: decimal.NewFromInt(1),
		},
		{
			name: "b valid",
			a:    decimal.NullDecimal{Decimal: decimal.NewFromInt(0), Valid: false},
			b:    decimal.NullDecimal{Decimal: decimal.NewFromInt(2), Valid: true},
			want: decimal.NewFromInt(2),
		},
		{
			name: "both invalid",
			a:    decimal.NullDecimal{Decimal: decimal.NewFromInt(0), Valid: false},
			b:    decimal.NullDecimal{Decimal: decimal.NewFromInt(0), Valid: false},
			want: decimal.NewFromInt(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.a, tt.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
