package task14

type CustomerType string

type (
	DiscountRate      float64
	DiscountThreshold float64
	DiscountAmount    float64
	ItemQuantity      int
)

const (
	CustomerVIP CustomerType = "VIP"
	CustomerNew CustomerType = "NEW"
)

const (
	vipDiscountRate           DiscountRate      = 0.1
	newDiscountRate           DiscountRate      = 0.05
	fixedDiscountThreshold    DiscountThreshold = 1000
	fixedDiscountAmount       DiscountAmount    = 50
	quantityDiscountThreshold ItemQuantity      = 10
	quantityDiscountRate      DiscountRate      = 0.01
)

var customerDiscountRates = map[CustomerType]DiscountRate{
	CustomerVIP: vipDiscountRate,
	CustomerNew: newDiscountRate,
}

type OrderService struct{}

// Calc возвращает стоимость заказа с учетом скидок.
// items содержит товары с их ценами и количеством.
// customerType задает тип клиента.
// Для VIP скидка составляет 10%. Для NEW скидка составляет 5%.
// Для остальных типов клиентов процентной скидки нет.
// Если сумма после скидки больше 1000, из нее вычитается 50.
// Затем применяется скидка 1%, если общее количество товаров больше 10.
// Если items пустой или равен nil, возвращается 0.
func (OrderService) Calc(items []Item, customerType CustomerType) float64 {
	total := calculateSubtotal(items)
	total = applyCustomerDiscount(total, customerType)
	total = applyFixedDiscount(total)
	return applyQuantityDiscount(total, calculateQuantity(items))
}

func calculateSubtotal(items []Item) float64 {
	var subtotal float64
	for _, item := range items {
		subtotal += item.Price() * float64(item.Quantity())
	}
	return subtotal
}

func calculateQuantity(items []Item) ItemQuantity {
	var quantity ItemQuantity
	for _, item := range items {
		quantity += ItemQuantity(item.Quantity())
	}
	return quantity
}

func applyCustomerDiscount(total float64, customerType CustomerType) float64 {
	discount := customerDiscountRates[customerType]
	return total * (1 - float64(discount))
}

func applyFixedDiscount(total float64) float64 {
	if total > float64(fixedDiscountThreshold) {
		return total - float64(fixedDiscountAmount)
	}
	return total
}

func applyQuantityDiscount(total float64, quantity ItemQuantity) float64 {
	if quantity > quantityDiscountThreshold {
		return total * (1 - float64(quantityDiscountRate))
	}
	return total
}
