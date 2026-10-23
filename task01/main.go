package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Println("Введите два числа и операцию (+, -, *, /):")
	var a, b float64
	var operation string
	read(&a, &b, &operation)
	var result float64
	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fail("деление на ноль")
		}
		result = a / b
	default:
		fail("неизвестная операция")
	}
	if math.IsNaN(result) || math.IsInf(result, 0) {
		fail("результат выходит за допустимые пределы")
	}
	fmt.Println("Результат:", result)
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
