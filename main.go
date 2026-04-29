

import "fmt"
package main

import (
"awesomeProject/internal/commission"
"fmt"
)

func main() {
	var amount int64
	var cardType int

	fmt.Println("=== Перевод средств ===")
	fmt.Print("Введите сумму перевода: ")
	fmt.Scan(&amount)

	fmt.Print("Тип карты (1 - Alif, 2 - Другая): ")
	fmt.Scan(&cardType)

	isAlif := cardType == 1

	comm, err := commission.Calculate(amount, isAlif)
	if err != nil {
		fmt.Printf("\n❌ Ошибка: %v\n", err)
		return
	}

	total := amount + comm


	fmt.Println("\n======= Alif Mobi =======")
	fmt.Println("Услуга:     Перевод")
	fmt.Printf("Сумма:      %d сум\n", amount)
	fmt.Printf("Комиссия:   %d сум\n", comm)
	fmt.Println("-------------------------")
	fmt.Printf("ИТОГО:      %d сум\n", total)
	fmt.Println("Статус:     Исполнено ✅")
	fmt.Println("=========================")
}

