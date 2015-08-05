package main

import "fmt"

func main(){
    m := make(map[string]int)
    m["Answer"] = 23
    fmt.Println("The value:",m["Answer"])

    m["answer"] = 46
    fmt.Println("The answer value:" , m["answer"])

    delete(m , "answer")
    fmt.Println("The value :" , m["answer"])

    v ,ok := m["answer"]
    if ok {
        fmt.Println("m[answer] is " , v)
    } else {
        fmt.Println("m[answer] is not been set")
    }
}
