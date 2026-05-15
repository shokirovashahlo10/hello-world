package main

import (
<<<<<<< HEAD
<<<<<<< HEAD
	"awesomeProject3/internal/commission"
=======
>>>>>>> c72f24fe2fed03af4ef9f7b0fac309fd8ac7a984
=======
>>>>>>> origin/main
	"fmt"
)

func main() {
<<<<<<< HEAD
<<<<<<< HEAD
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
=======
=======
>>>>>>> origin/main
	fmt.Println("Hello Alif")
	var name string = "iphone 13 pro"
	var brand string = "apple "
	var price int64 = 120000
	var isAvailable bool = true

	paymente := float64(price) / 12
	fmt.Println("==========Alif Shop========")
	fmt.Printf("товар: %s\n", name)
	fmt.Printf("бренд: %s\n", brand)
	fmt.Printf("цена: %d\n", price)
	fmt.Printf("в наличии: %t\n", isAvailable)
	fmt.Printf("рассрочка : 12 мес → %.0f сум/мес\n", paymente)
	fmt.Println("===================")

<<<<<<< HEAD
>>>>>>> c72f24fe2fed03af4ef9f7b0fac309fd8ac7a984
=======
>>>>>>> origin/main
}
