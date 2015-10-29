package main

import "fmt"

func main(){
    var a[2] string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[0] , a[1])

    fmt.Println(a)

    //array init
    var x [5]int
    x[0] = 1
    x[1] = 2
    for _,i := range x {
        fmt.Println(i)
    }

    //array init 2
    var y = [5]int{1,2}
    fmt.Println(y)
    
    //array init 3
    var z = [...]bool{true , false , true , false}
    fmt.Println(z,len(z))
    
    //array init 4
    var j = []string{
        "Monday",
        "Tuesday",
        "Wednesday",
        "Thursday",
        "Friday",
        "Saturday",
        "Sunday",       //,必须，或写成"Saturday"}
    }
    fmt.Println(j)
}
