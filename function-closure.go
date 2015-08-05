package main

import "fmt"

func closure() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    pos , neg := closure(),closure()

    for i:=0;i<10;i++ {
        fmt.Println(pos(i) , neg(-2*i))
    }
}
