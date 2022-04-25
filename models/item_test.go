package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateFinalPrice(t *testing.T) {
	var items = []Item{
		{
			Name:       "Pen",
			Price:      10,
			Quantity:   1,
			Type:       "imported",
			Tax:        1.47,
			FinalPrice: 11.47,
		},
		{
			Name:       "wood",
			Price:      150,
			Quantity:   1,
			Type:       "raw",
			Tax:        18.75,
			FinalPrice: 168.75,
		},
		{
			Name:       "furniture",
			Price:      1000,
			Quantity:   10,
			Type:       "manufactured",
			Tax:        15,
			FinalPrice: 10150,
		},
		{
			Name:       "chocolates",
			Price:      2000,
			Quantity:   5,
			Type:       "imported",
			Tax:        295,
			FinalPrice: 11475,
		},
	}
	for _, item := range items {
		validItem := NewItem(item.Name, item.Price, item.Quantity, item.Type)
		validItem.CalculateFinalPrice()
		require.Equal(t, validItem.FinalPrice, item.FinalPrice)
		require.Equal(t, validItem.Tax, item.Tax)
	}
}
