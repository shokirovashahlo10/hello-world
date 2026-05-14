package main

import (
	"fmt"

	"./internal/commission"
)

func main() {
	var sName, sLast string // Отправитель
	var rName, rLast string // Получатель
	var card string
	var amount int64

	fmt.Println("Введите данные:")
	fmt.Print("Имя Фамилия отправителя: ")
	fmt.Scan(&sName, &sLast)

	fmt.Print("Имя Фамилия получателя: ")
	fmt.Scan(&rName, &rLast)

	fmt.Print("Номер карты (16 цифр): ")
	fmt.Scan(&card)

	fmt.Print("Сумма: ")
	fmt.Scan(&amount)

	// Печатаем чек
	fmt.Println("\n======= Alif Mobi =======")
	fmt.Printf("Отправитель: %s %s\n", sLast, sName)
	fmt.Printf("Получатель:  %s %s\n", rLast, rName)
	fmt.Printf("Карта:       %s\n", commission.MaskCard(card))
	fmt.Println("-------------------------")
	fmt.Printf("Сумма:       %d сум\n", amount)
	fmt.Printf("Транзакция:  №%d\n", commission.GenID())
	fmt.Printf("Дата:        %s\n", commission.GetTime())
	fmt.Println("Статус:      Исполнено ✅")
	fmt.Println("=========================")
}
