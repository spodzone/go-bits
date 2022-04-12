package main

import (
	"fmt"
	"math"
)

func primes(x int64) []int64 {
	var ret = []int64{}
	for t := int64(2); t < x; t++ {
		if isprime(t) {
			ret = append(ret, t)
		}
	}
	return ret
}

func isprime(x int64) bool {
	limit := int64(math.Floor(math.Sqrt(float64(x))) + 1)
	// fmt.Printf("ISPRIME: %d up to %d\n", x, limit)
	var i int64
	ps := primes(limit)
	for _, i = range ps {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Hello world!")
	fmt.Println("Test of isprime:")
	var lst = []int64{2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("lst=", lst)
	for _, x := range lst {
		fmt.Printf("Test isprime %d=%t\n", x, isprime(x))
	}
	fmt.Println("Primes up to 100", primes(int64(100)))
}
