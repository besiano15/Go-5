package main

import "fmt"

func main(){
    sl := make([]int , 10)

    for i := range sl {
        sl[i] = 1 << uint(i)
    }

    for _ , val := range sl {
        fmt.Printf("%d\n" , val)
    }
}
