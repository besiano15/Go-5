package main

import (
    "fmt"
)

func main(){
    var a string = "hello" 
    var b string = "world"

    fmt.Println("a length:",len(a) , " b length:",len(b),"\n")
    fmt.Println(a[2],"\n")

    fmt.Println(a+b)
}
