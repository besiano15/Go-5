package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt struct {
	Num float64
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Error %v \n", e.Num)
}

func Sqrt(x float64) (res float64, err error) {
	if x < 0 {
		err = &ErrNegativeSqrt{x}
		return
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
