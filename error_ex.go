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
    res , err := Sqrt(2)
    if err == nil {
        fmt.Println(res)
    }

    res1,err1 := Sqrt(-2)
    if err1 != nil {
        fmt.Println(err1)
    } else {
        fmt.Println(res1)
    }
}
