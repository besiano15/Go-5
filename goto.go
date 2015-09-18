package main

import "fmt"

func main() {
    MyFunc()
}

func MyFunc() {
    i := 0
    LooP :
    fmt.Println(i)
    i++
    if i<10 {
        goto LooP
    }
}
