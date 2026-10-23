package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите целое число для таблицы умножения:")
	var n int
	read(&n)
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1
	if n > maxInt/10 || n < minInt/10 {
		fail("слишком большое число для таблицы")
	}
	printTable(n)
}

func printTable(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", n, i, n*i)
	}
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
