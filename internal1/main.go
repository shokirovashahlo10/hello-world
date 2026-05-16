package main

import "awesomeProject3/internal1/product"

func main() {
	phone := product.Product{
		ID:     1,
		Name:   "Xiaomi Redmi Note 14 6/128 ГБ",
		Price:  2499000,
		Months: 12,
	}

	product.PrintProduct(phone)
}
