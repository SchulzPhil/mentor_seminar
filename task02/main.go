package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите целое число:")
	var n int
	read(&n)
	if n > 0 {
		fmt.Println("Положительное")
	} else if n < 0 {
		fmt.Println("Отрицательное")
	} else {
		fmt.Println("Ноль")
	}
	if n%2 == 0 {
		fmt.Println("Чётное")
	} else {
		fmt.Println("Нечётное")
	}
	fmt.Println("Делится на 3:", n%3 == 0)
	fmt.Println("Делится на 5:", n%5 == 0)
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
