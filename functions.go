package main

import (
    "fmt"
    "math"
)

func main(){
    fun := func(x,y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }

    fmt.Println(fun(6,8))
}
