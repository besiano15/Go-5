package main

import (
    "fmt"
    "time"
)

func main(){
    fmt.Println("Welcome to the playground")
    _,month,day := time.Now().Date()
    fmt.Println("Now time is",month ,day)

    for i:=0 ; i<10 ; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println("Hello")
    }
}
