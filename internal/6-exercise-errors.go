package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if (x < 0) {
		return 0.0, ErrNegativeSqrt(x)
	}
	
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
	return z, nil;
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
