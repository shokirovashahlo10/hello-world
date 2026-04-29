package commission

import "errors"

func Calculate(amount int64, isAlif bool) (int64, error) {
	if amount < 500 {
		return 0, errors.New("минимальная сумма перевода 500 сум")
	}
	if amount > 15000000 {
		return 0, errors.New("максимальная сумма перевода 15 000 000 сум")
	}

	if isAlif {
		return 0, nil
	}

	commission := (amount * 29) / 10000
	return commission, nil
}
