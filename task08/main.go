package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите целое число:")
	var n int
	read(&n)
	fmt.Println("Простое:", isPrime(n))
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= n/i; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
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
