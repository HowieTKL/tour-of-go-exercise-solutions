package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	const DeltaThreshold = 0
	z := 1.0
	delta := 1.0
	fmt.Printf("DeltaThreshold: %v\n", DeltaThreshold)
	for i := 0; delta > DeltaThreshold; i++ {
		prev := z;
		z -= (z*z - x) / (2*z)
		delta = math.Abs(prev - z);
		fmt.Printf("%v %v delta=%v\n", i, z, delta)
	}
	return z;
}

func main() {
	fmt.Println(Sqrt(2))
}
