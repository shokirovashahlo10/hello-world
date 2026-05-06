package commission

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Calculate(amount int64, isAlif bool) (int64, error) {
	if amount < 500 {
		return 0, errors.New("минимум 500 сум")
	}
	if amount > 15000000 {
		return 0, errors.New("максимум 15 млн сум")
	}
	if isAlif {
		return 0, nil
	}
	return (amount * 29) / 10000, nil
}

func MaskCard(number string) string {
	if len(number) < 4 {
		return "UZCARD**0000"
	}
	last := number[len(number)-4:]
	return fmt.Sprintf("UZCARD**%s", last)
}

func GenID() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(900000000) + 100000000
}

func GetTime() string {
	return time.Now().Format("02.01.2006 15:04")
}
