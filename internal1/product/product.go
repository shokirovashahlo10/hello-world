package product

import "fmt"

type ProductID int
type Price int
type Months int

type Product struct {
	ID     ProductID
	Name   string
	Price  Price
	Months Months
}

func MonthlyPayment(p Product) int {

	if p.Months < 3 || p.Months > 24 {
		return 0
	}
	return int(p.Price) / int(p.Months)
}

func PrintProduct(p Product) {
	payment := MonthlyPayment(p)

	fmt.Println("===== Alif Shop =====")
	fmt.Printf("ID:      %d\n", p.ID)
	fmt.Printf("Товар:   %s\n", p.Name)
	fmt.Printf("Цена:    %d сум\n", p.Price)
	fmt.Printf("Срок:    %d мес.\n", p.Months)
	if payment > 0 {
		fmt.Printf("Платеж:  %d сум/мес\n", payment)
	} else {
		fmt.Println("Ошибка:  Неверный срок рассрочки")
	}
	fmt.Println("=====================")
}
