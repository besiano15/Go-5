package main

import (
    "fmt"
)

func main() {
    var x = [5]int{1,2,3,4}
    x[4] = 5 

    var sum int
    for _ , i := range x {
        sum += i
    }

    fmt.Println(sum)
}
