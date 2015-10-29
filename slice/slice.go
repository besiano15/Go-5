package main

import (
    "fmt"
)

func main() {
    var x = make([]float64 ,5 )
    fmt.Println("Capcity:",cap(x) , "Length:",len(x))

    var y = make([]float64, 5, 10)
    fmt.Println("Capcity:",cap(y) , "Length:",len(y))

    for i:=0 ;i<len(x);i++ {
        x[i] = float64(i)
    }

    fmt.Println(x)

    for i:=0;i<len(y);i++ {
        y[i] = float64(i)
    }

    fmt.Println(y)
    
}
