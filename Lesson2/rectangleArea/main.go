package main

import (
	"fmt"
)

func main()  {
	var distance1, distance2 float64

	fmt.Println("Введите первый размер прямоугольника")
	fmt.Scanln(&distance1)

	fmt.Println("Введите второй размер прямоугольника")
	fmt.Scanln(&distance2)

	fmt.Printf("Площадь прямоугольника = %f \n", distance1 * distance2)
}
