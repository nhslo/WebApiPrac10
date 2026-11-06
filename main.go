package main

import "fmt"

type Sale struct {
	ID                int
	Product, Category string
	Price             float64
	Quantity          int
}

// Выручка каждой продажи — цена, умноженная на количество.
func CalculateRevenue(sales []Sale) float64 {
	sum := 0.0
	for _, sale := range sales {
		sum = sum + sale.Price*float64(sale.Quantity)
	}
	return sum
}

func GetTopProduct(sales []Sale) (string, int) {
	counts := map[string]int{}
	for _, sale := range sales {
		counts[sale.Product] = counts[sale.Product] + sale.Quantity
	}
	best := ""
	count := 0
	for name, quantity := range counts {
		if quantity > count {
			best = name
			count = quantity
		}
	}
	return best, count
}

func RevenueByCategory(sales []Sale) map[string]float64 {
	result := map[string]float64{}
	for _, sale := range sales {
		result[sale.Category] = result[sale.Category] + sale.Price*float64(sale.Quantity)
	}
	return result
}

// Через указатель меняем цену выбранной продажи.
func ApplyDiscount(sale *Sale, discount float64) {
	discountAmount := sale.Price * discount / 100
	sale.Price = sale.Price - discountAmount
}

func main() {
	sales := []Sale{
		{ID: 1, Product: "Ноутбук", Category: "Электроника", Price: 350000, Quantity: 2},
		{ID: 2, Product: "Мышь", Category: "Аксессуары", Price: 5000, Quantity: 10},
		{ID: 3, Product: "Монитор", Category: "Электроника", Price: 85000, Quantity: 3},
		{ID: 4, Product: "Клавиатура", Category: "Аксессуары", Price: 12000, Quantity: 5},
		{ID: 5, Product: "Наушники", Category: "Аксессуары", Price: 25000, Quantity: 4},
		{ID: 6, Product: "Ноутбук", Category: "Электроника", Price: 350000, Quantity: 1},
	}
	discounts := [3]float64{5, 10, 15}
	count := 0
	fmt.Println("Продажи:")
	for _, sale := range sales {
		fmt.Printf("%d. %s | %s | %.0f тг | %d шт. | сумма %.0f тг\n", sale.ID, sale.Product, sale.Category, sale.Price, sale.Quantity, sale.Price*float64(sale.Quantity))
		count = count + sale.Quantity
	}
	total := CalculateRevenue(sales)
	best, quantity := GetTopProduct(sales)
	averageSale := total / float64(len(sales))
	fmt.Println("\nЗаписей:", len(sales))
	fmt.Println("Продано единиц:", count)
	fmt.Printf("Выручка: %.0f тг\n", total)
	fmt.Printf("Средняя сумма продажи: %.0f тг\n", averageSale)
	fmt.Println("Самый продаваемый товар:", best, "—", quantity, "шт.")
	fmt.Println("Выручка по категориям:")
	for name, revenue := range RevenueByCategory(sales) {
		fmt.Printf("%s: %.0f тг\n", name, revenue)
	}
	old := sales[0].Price
	ApplyDiscount(&sales[0], discounts[1])
	fmt.Printf("Скидка %.0f%%: %s, %.0f → %.0f тг\nВыручка после скидки: %.0f тг\n", discounts[1], sales[0].Product, old, sales[0].Price, CalculateRevenue(sales))
}
