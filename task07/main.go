package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите N (неотрицательное целое число):")
	var n int
	read(&n)
	if n < 0 {
		fail("N не должно быть отрицательным")
	}
	var sum float64
	for i := 0; i < n; i++ {
		value := i + 1
		if value%2 == 0 && value%3 == 0 {
			sum += float64(value)
		}
	}
	fmt.Println("Сумма:", sum)
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
