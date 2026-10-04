package task14_test

import (
	"math"
	"slices"
	"testing"

	task14 "github.com/gregory-vc/algos_2026/task_14"
)

func TestOrderServiceCalc(t *testing.T) {
	basket := []task14.Item{
		task14.NewItem("Книга", 100, 2),
		task14.NewItem("Игра", 200, 1),
	}
	tests := []struct {
		name         string
		items        []task14.Item
		customerType task14.CustomerType
		want         float64
	}{
		{
			name: "nil_basket", customerType: "REGULAR", want: 0,
		},
		{
			name: "nil_basket_vip", customerType: task14.CustomerVIP, want: 0,
		},
		{
			name: "nil_basket_new", customerType: task14.CustomerNew, want: 0,
		},
		{
			name: "empty_basket", items: []task14.Item{}, customerType: "REGULAR", want: 0,
		},
		{
			name: "regular_customer", items: basket, customerType: "REGULAR", want: 400,
		},
		{
			name: "vip_customer", items: basket, customerType: task14.CustomerVIP, want: 360,
		},
		{
			name: "new_customer", items: basket, customerType: task14.CustomerNew, want: 380,
		},
		{
			name: "unknown_customer", items: basket, customerType: "UNKNOWN", want: 400,
		},
		{
			name: "empty_customer_type", items: basket, customerType: "", want: 400,
		},
		{
			name: "customer_type_is_case_sensitive", items: basket, customerType: "vip", want: 400,
		},
		{
			name: "fractional_prices",
			items: []task14.Item{
				task14.NewItem("Тетрадь", 19.99, 3),
				task14.NewItem("Ручка", 0.1, 2),
			},
			customerType: "REGULAR", want: 60.17,
		},
		{
			name: "zero_quantity",
			items: []task14.Item{
				task14.NewItem("Книга", 100, 0),
				task14.NewItem("Игра", 200, 1),
			},
			customerType: "REGULAR", want: 200,
		},
		{
			name: "free_item", items: []task14.Item{task14.NewItem("Подарок", 0, 3)},
			customerType: "REGULAR", want: 0,
		},
		{
			name: "below_fixed_discount_threshold", items: []task14.Item{task14.NewItem("Товар", 999.99, 1)},
			customerType: "REGULAR", want: 999.99,
		},
		{
			name: "at_fixed_discount_threshold", items: []task14.Item{task14.NewItem("Товар", 1000, 1)},
			customerType: "REGULAR", want: 1000,
		},
		{
			name: "just_above_fixed_discount_threshold", items: []task14.Item{task14.NewItem("Товар", 1000.01, 1)},
			customerType: "REGULAR", want: 950.01,
		},
		{
			name: "fixed_discount_applied_once_to_entire_basket",
			items: []task14.Item{
				task14.NewItem("Книга", 300, 2),
				task14.NewItem("Игра", 300, 2),
			},
			customerType: "REGULAR", want: 1150,
		},
		{
			name: "unknown_customer_still_gets_fixed_discount", items: []task14.Item{task14.NewItem("Товар", 1200, 1)},
			customerType: "UNKNOWN", want: 1150,
		},
		{
			name: "vip_discount_reduces_total_below_threshold", items: []task14.Item{task14.NewItem("Товар", 1100, 1)},
			customerType: task14.CustomerVIP, want: 990,
		},
		{
			name: "new_discount_reduces_total_below_threshold", items: []task14.Item{task14.NewItem("Товар", 1050, 1)},
			customerType: task14.CustomerNew, want: 997.5,
		},
		{
			name: "vip_then_fixed_discount", items: []task14.Item{task14.NewItem("Товар", 1200, 1)},
			customerType: task14.CustomerVIP, want: 1030,
		},
		{
			name: "new_then_fixed_discount", items: []task14.Item{task14.NewItem("Товар", 1200, 1)},
			customerType: task14.CustomerNew, want: 1090,
		},
	}

	service := task14.OrderService{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := service.Calc(tt.items, tt.customerType)
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("Calc() = %g, want %g", got, tt.want)
			}
		})
	}
}

func TestOrderServiceCalcQuantityDiscount(t *testing.T) {
	basketWithEmptyPositions := make([]task14.Item, 11)
	basketWithEmptyPositions[0] = task14.NewItem("Книга", 50, 10)

	tests := []struct {
		name         string
		items        []task14.Item
		customerType task14.CustomerType
		want         float64
	}{
		{
			name:         "ten_units_without_quantity_discount",
			items:        []task14.Item{task14.NewItem("Книга", 50, 10)},
			customerType: "REGULAR", want: 500,
		},
		{
			name:         "eleven_units_in_one_position",
			items:        []task14.Item{task14.NewItem("Книга", 50, 11)},
			customerType: "REGULAR", want: 544.5,
		},
		{
			name: "eleven_units_in_two_positions",
			items: []task14.Item{
				task14.NewItem("Книга", 50, 5),
				task14.NewItem("Игра", 50, 6),
			},
			customerType: "REGULAR", want: 544.5,
		},
		{
			name:         "eleven_positions_but_ten_units",
			items:        basketWithEmptyPositions,
			customerType: "REGULAR", want: 500,
		},
		{
			name:         "ten_units_with_fixed_discount",
			items:        []task14.Item{task14.NewItem("Книга", 200, 10)},
			customerType: "REGULAR", want: 1950,
		},
		{
			name:         "ten_units_with_vip_and_fixed_discount",
			items:        []task14.Item{task14.NewItem("Книга", 200, 10)},
			customerType: task14.CustomerVIP, want: 1750,
		},
		{
			name:         "ten_units_with_new_and_fixed_discount",
			items:        []task14.Item{task14.NewItem("Книга", 200, 10)},
			customerType: task14.CustomerNew, want: 1850,
		},
		{
			name:         "quantity_discount_after_fixed_discount",
			items:        []task14.Item{task14.NewItem("Книга", 100, 20)},
			customerType: "REGULAR", want: 1930.5,
		},
		{
			name:         "quantity_discount_after_vip_and_fixed_discount",
			items:        []task14.Item{task14.NewItem("Книга", 100, 20)},
			customerType: task14.CustomerVIP, want: 1732.5,
		},
		{
			name:         "quantity_discount_after_new_and_fixed_discount",
			items:        []task14.Item{task14.NewItem("Книга", 100, 20)},
			customerType: task14.CustomerNew, want: 1831.5,
		},
		{
			name:         "unknown_customer_gets_quantity_discount",
			items:        []task14.Item{task14.NewItem("Книга", 50, 11)},
			customerType: "UNKNOWN", want: 544.5,
		},
	}

	service := task14.OrderService{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.Calc(tt.items, tt.customerType); got != tt.want {
				t.Errorf("Calc() = %g, want %g", got, tt.want)
			}
		})
	}
}

func TestOrderServiceCalcDoesNotModifyItems(t *testing.T) {
	items := []task14.Item{
		task14.NewItem("Книга", 300, 2),
		task14.NewItem("Игра", 300, 2),
	}
	original := slices.Clone(items)
	service := task14.OrderService{}

	for _, customerType := range []task14.CustomerType{"REGULAR", task14.CustomerVIP, task14.CustomerNew} {
		service.Calc(items, customerType)
		if !slices.Equal(items, original) {
			t.Fatalf("Calc() changed items for customer type %q", customerType)
		}
	}
}
