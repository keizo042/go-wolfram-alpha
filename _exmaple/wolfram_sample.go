package main

import (
    "fmt"
    "github.com/KeizoBookman/go-wolfram-alpha"
    "github.com/KeizoBookman/go-wolfram-alpha/mathematica/"
)

func main() {
    c: = mathematica.New()
    c.Fetch("1 + 1")
    
}
