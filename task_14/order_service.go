package task14

type OrderService struct{}

// Calc возвращает стоимость заказа с учетом скидок.
// items содержит товары с их ценами и количеством.
// customerType задает тип клиента.
// Для VIP скидка составляет 10%. Для NEW скидка составляет 5%.
// Для остальных типов клиентов процентной скидки нет.
// Если сумма после скидки больше 1000, из нее вычитается 50.
// Если items пустой или равен nil, возвращается 0.
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
