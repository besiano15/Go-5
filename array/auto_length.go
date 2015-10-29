package main

import (
    "fmt"
)

func main() {
    var x = [...]string{
        "Monday",
        "Tuesday",
        "Wednesday",
        "Thursday",
        "Friday",
        "Saturday",
        "Sunday",   //逗号必须的,可以把结束大括号提上来
    }

    for _, day := range x {
        fmt.Println(day)
    }
}
