package main

import "fmt"

func main(){
    var sum = 0
    for i:=1 ;i<=100;i++ {
        sum += i
        fmt.Println(i);
    }

    fmt.Println("result is:",sum);

}
