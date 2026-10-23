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
	if n < 0 {
		fail("N не должно быть отрицательным")
	}
	positive, negative, zero := 0, 0, 0
	for i := 0; i < n; i++ {
		var value float64
		read(&value)
		if value > 0 {
			positive++
		} else if value < 0 {
			negative++
		} else {
			zero++
		}
	}
	fmt.Println("Положительных:", positive)
	fmt.Println("Отрицательных:", negative)
	fmt.Println("Нулей:", zero)
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
