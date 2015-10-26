package main

import (
    "fmt"
)

func main() {
    //var a , b ,c := false , true , "java" 
    var x string = "hello world" //显式变量定义
    var y string
    y = "hello too" // 定义 赋值分离

    var z = "hello either" //类型推导

    q := "hello pg" //自动类型推导，不需要var，但必须在func内
    a , b ,c := false , true , "java" 
    fmt.Println(a,b,c)

    fmt.Println(x,y,z,q)

}
