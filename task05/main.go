package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Println("Введите количество чисел N, затем N чисел:")
	var n int
	read(&n)
	if n <= 0 {
		fail("N должно быть положительным")
	}
	var value float64
	read(&value)
	minimum, maximum := value, value
	for i := 1; i < n; i++ {
		read(&value)
		if value < minimum {
			minimum = value
		}
		if value > maximum {
			maximum = value
		}
	}
	fmt.Println("Максимум:", maximum)
	fmt.Println("Минимум:", minimum)
}

func read(values ...any) {
	if _, err := fmt.Fscan(os.Stdin, values...); err != nil {
		fail("некорректный или неполный ввод")
	}
	for _, value := range values {
		if number, ok := value.(*float64); ok && (math.IsNaN(*number) || math.IsInf(*number, 0)) {
			fail("ожидалось конечное число")
		}
	}
}

func fail(message string) {
	fmt.Fprintln(os.Stderr, "Ошибка:", message)
	os.Exit(1)
}
