package main

import "fmt"

func main(){
    sl := []int{2,4,5,7,9,12}

    fmt.Println("sl = ",sl)

    fmt.Println("sl[1:4] =",sl[1:4])

    
    fmt.Println("sl[:5] =",sl[:5])
    
    fmt.Println("sl[1:] =",sl[1:])
}
