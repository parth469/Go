package main

import "fmt"

const name = "test"

func main() {
    const a = 8
    fmt.Println(a)
    testGlobal()
}

func testGlobal() {
    fmt.Println(name)
    //The below line will give compiler error as a is a local constant
    //fmt.Println(a)
}