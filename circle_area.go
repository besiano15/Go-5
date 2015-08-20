package main

import (
    "fmt"
    "math"
)
//变量类型不可推导，初始必须赋值，切后续不可变更
func main() {
    var radius float64 = 10
    var area = math.Pow(radius , 2) * math.Pi
    fmt.Println(area)
}
