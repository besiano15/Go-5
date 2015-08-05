package main

import "fmt"

func fabnacci() func() int {
    sl := []int{0,1}
    j := 0
    return func () int {
        if j>1{
            sl = append(sl , sl[j-2] + sl[j-1])
        }
        j += 1
        return sl[ j-1 ]
    }
}

func main(){
    f := fabnacci()

    for i:=0 ;i<10;i++ {
        fmt.Println(f())
    }
}
