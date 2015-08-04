package main

import "fmt"

func main(){
    s :=[]int{2,3,56,6,87,8}
    fmt.Println("s == " , s)

    for i := 0;i<len(s) ;i++ {
        fmt.Printf("s[%d] == %d\n" , i , s[i])
    }
}
