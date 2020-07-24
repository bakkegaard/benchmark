package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i < int(math.Sqrt(float64(n)))+1; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Requires one argument, remember target")
	}
	var sum int64 = 0
	var target, _ = strconv.ParseInt(os.Args[1], 10, 64)

	var target2 = int(target)

	for i := 1; i < target2; i++ {
		if isPrime(i) {
			sum++
		}
	}
	fmt.Println(sum)
}
