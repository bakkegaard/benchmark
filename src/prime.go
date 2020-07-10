package main

import "os"
import "fmt"
import "strconv"
import "math"

func isPrime(n int) bool {
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
	var target, error = strconv.ParseInt(os.Args[1], 10, 64)

	//Stupid...
	_ = error

	var target2 = int(target)

	for i := 2; i < target2; i++ {
		if isPrime(i) {
			sum++
		}
	}
	fmt.Println(sum)
}
