package main

import (
    "fmt"
    "github.com/maxb/go-dep-version-test/alpha"
    "github.com/maxb/go-dep-version-test/beta"
)

func main() {
    fmt.Println("Hello, world!")
    fmt.Println(alpha.Alpha())
    fmt.Println(beta.Beta())
}
