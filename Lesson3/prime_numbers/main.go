package main

import (
	"fmt"
	"math"
)

//isPrimeNumber
func isPrimeNumber(number int64) bool {
	var x, y, n int64
	nsqrt := math.Sqrt(float64(number))

	is_prime := make([]bool, number + 2)// []bool{}1

	for x = 1; float64(x) <= nsqrt; x++ {
		for y = 1; float64(y) <= nsqrt; y++ {
			n = 4*(x*x) + y*y
			if n <= number && (n%12 == 1 || n%12 == 5) {
				is_prime[n] = !is_prime[n]
			}
			n = 3*(x*x) + y*y
			if n <= number && n%12 == 7 {
				is_prime[n] = !is_prime[n]
			}
			n = 3*(x*x) - y*y
			if x > y && n <= number && n%12 == 11 {
				is_prime[n] = !is_prime[n]
			}
		}
	}

	for n = 5; float64(n) <= nsqrt; n++ {
		if is_prime[n] {
			for y = n * n; y < number; y += n * n {
				is_prime[y] = false
			}
		}
	}

	is_prime[2] = true
	is_prime[3] = true

	//primes := make([]int64, 0, 1270606)
	for x = 0; int64(x) < int64(len(is_prime)-1); x++ {
		if is_prime[x] && x == number {
			return true
		}
	}

	return false

}

func main() {

	var number int64

	fmt.Println("Введите число")
	fmt.Scanln(&number)

	if isPrimeNumber(number) {
		fmt.Println("Это число простое")
	} else {
		fmt.Println("Это число составное")
	}



}
