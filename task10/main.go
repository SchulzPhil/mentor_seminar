package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите количество чисел N, затем N целых чисел:")
	var n int
	read(&n)
	if n <= 0 {
		fail("N должно быть положительным")
	}
	var sum float64
	var minimum, maximum int
	even, odd := 0, 0
	for i := 0; i < n; i++ {
		var value int
		read(&value)
		sum += float64(value)
		if i == 0 || value < minimum {
			minimum = value
		}
		if i == 0 || value > maximum {
			maximum = value
		}
		if value%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	fmt.Println("Сумма:", sum)
	fmt.Println("Среднее:", sum/float64(n))
	fmt.Println("Максимум:", maximum)
	fmt.Println("Минимум:", minimum)
	fmt.Println("Чётных:", even)
	fmt.Println("Нечётных:", odd)
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
