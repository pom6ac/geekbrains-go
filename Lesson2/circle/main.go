package main

import (
	"fmt"
	"math"
)

func main()  {
	var square float64
	var diameter float64

	fmt.Println("Введите плозадь окружности")
	fmt.Scanln(&square)

	diameter = math.Sqrt(square / math.Pi)

	fmt.Printf("Диаметр окружности равен %f\n", diameter)
	fmt.Printf("Длина окружности равна %f\n", diameter * math.Pi)
}
