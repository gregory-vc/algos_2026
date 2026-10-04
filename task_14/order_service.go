package task14

type OrderService struct{}

func (OrderService) Calc(items []Item, customerType string) float64 {
	var s float64
	for _, i := range items {
		s += i.Price() * float64(i.Quantity())
	}

	if customerType == "VIP" {
		s = s * 0.9
	}

	if customerType == "NEW" {
		s = s * 0.95
	}

	if s > 1000 {
		s = s - 50
	}

	return s
}
