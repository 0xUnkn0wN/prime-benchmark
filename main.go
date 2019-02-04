package main

import (
	"fmt"
)

func main() {
	max := uint(100000)

	fmt.Println(calcPrime(max))
}

func calcPrime(max uint) (returnVal []uint) {
	for i := uint(2); i <= max; i++ {
		// TODO: add multi core support
		if isPrime(i) {
			returnVal = append(returnVal, i)
		}
	}

	return
}

func isPrime(number uint) bool {
	for a := uint(2); a <= number/2; a++ {
		if number%a == 0 {
			return false
		}
	}

	return true
}
