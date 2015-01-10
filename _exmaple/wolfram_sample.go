package main

import (
    "fmt"
    "github.com/KeizoBookman/go-wolfram-alpha"
    "github.com/KeizoBookman/go-wolfram-alpha/mathematica/"
)

func main() {
    c: = mathematica.New()
    q,err :=c.Fetch("1 +1")
    if err != nil {
	return -1
    }
}
