package task14_test

import (
	"testing"

	task14 "github.com/gregory-vc/algos_2026/task_14"
)

func TestNewItem(t *testing.T) {
	tests := []struct {
		name     string
		itemName string
		price    float64
		quantity int
	}{
		{name: "product", itemName: "Книга", price: 123.45, quantity: 3},
		{name: "zero_values", itemName: "", price: 0, quantity: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := task14.NewItem(tt.itemName, tt.price, tt.quantity)
			if got := item.Name(); got != tt.itemName {
				t.Errorf("Name() = %q, want %q", got, tt.itemName)
			}
			if got := item.Price(); got != tt.price {
				t.Errorf("Price() = %g, want %g", got, tt.price)
			}
			if got := item.Quantity(); got != tt.quantity {
				t.Errorf("Quantity() = %d, want %d", got, tt.quantity)
			}
		})
	}
}
