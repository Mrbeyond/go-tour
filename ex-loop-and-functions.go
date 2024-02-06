package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {
	z = float64(1)
	// z = x / 2.0
	// z = x
	previousZ := z
	minuteChangeValue := 1e-15

	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Value of z is %v for iteration %v when x is %v \n", z, i, x)

		if math.Abs(previousZ-z) < minuteChangeValue {
			return
		}
		previousZ = z
	}
	return
}

func main() {
	for i := float64(1); i < 1025; i += i {
		fmt.Println(Sqrt(i), math.Sqrt(i), "\n ")
	}

}
