package main

import "fmt"

func main(){
    var sl []int
    fmt.Println(sl , len(sl) ,cap(sl))
    if sl == nil {
        fmt.Println("nil")
    }
}
