package main

import (
    "fmt"
    "runtime"
)

var city string = "Lxx"
var x int = 100

func main() {
    fmt.Printf("%s\n", runtime.Version())
    fmt.Printf("%s\n", city)
}
