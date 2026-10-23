package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите N (положительное целое число):")
	var n int
	read(&n)
	if n <= 0 {
		fail("N должно быть положительным")
	}
	var sum float64
	for i := 0; i < n; i++ {
		sum += float64(i) + 1
	}
	fmt.Println("Сумма:", sum)
	fmt.Println("Среднее:", sum/float64(n))
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
