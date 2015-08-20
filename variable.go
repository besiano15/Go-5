package main

import "fmt"

var c,python,java bool

func main() {
    //定义赋值
    var i int = 10

    //定以后赋值
    var j int 
    j = 20
    
    //类型推导
    var k = 30 
    //类型推导，仅在func内部，放在外面编译报错
    l := 40
    fmt.Println(i , c ,python,java)
    fmt.Println(j ,k ,l)
}
