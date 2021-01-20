package main

import "fmt"

func main()  {
	var number int64 = 347

	fmt.Println("Введите трехзначное число")
	//fmt.Scanln(&number)

	a := number / 100
	b := number / 10 % 10
	c := number % 10

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
