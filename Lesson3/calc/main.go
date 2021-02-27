package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var a, b, res float32
	var op string

	fmt.Print("Введите первое число: ")
	n, err := fmt.Scanln(&a)

	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}
	
	fmt.Print("Введите второе число: ")
	n, err := fmt.Scanln(&b)

	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	fmt.Print("Введите арифметическую операцию (+, -, *, /): ")
	n, err := fmt.Scanln(&op)

	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	fmt.Printf("Результат выполнения операции: %f\n", res)
}
