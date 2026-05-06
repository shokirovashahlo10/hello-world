package product

import "fmt"

func ExampleMonthlyPayment() {
	p := Product{
		ID:     101,
		Name:   "Xiaomi Redmi Note 14 6/128 ГБ",
		Price:  2499000,
		Months: 12,
	}

	payment := MonthlyPayment(p)
	fmt.Println(payment)

}
