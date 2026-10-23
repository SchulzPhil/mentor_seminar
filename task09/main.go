package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите баланс в целых копейках:")
	var balance int64
	read(&balance)
	if balance < 0 {
		fail("баланс не может быть отрицательным")
	}
	fmt.Println("Выберите действие: 1 — баланс, 2 — пополнить, 3 — снять")
	var action int
	read(&action)
	switch action {
	case 1:
	case 2, 3:
		fmt.Println("Введите сумму в целых копейках:")
		var amount int64
		read(&amount)
		if amount <= 0 {
			fail("сумма должна быть положительной")
		}
		if action == 2 {
			const maxInt64 = int64(1<<63 - 1)
			if amount > maxInt64-balance {
				fail("слишком большая сумма пополнения")
			}
			balance += amount
		} else {
			if amount > balance {
				fail("недостаточно средств")
			}
			balance -= amount
		}
	default:
		fail("неизвестное действие")
	}
	fmt.Printf("Баланс: %d.%02d руб.\n", balance/100, balance%100)
}

func read(values ...any) {
	if _, err := fmt.Fscan(os.Stdin, values...); err != nil {
		fail("некорректный или неполный ввод")
	}

}

func fail(message string) {
	fmt.Fprintln(os.Stderr, "Ошибка:", message)
	os.Exit(1)
}
