package main

import (
	"fmt"
	wolfram "github.com/KeizoBookman/go-wolfram-alpha"
)

func main() {
	id := ""
	c := wolfram.New(id)
	err := c.Get("oosaka")
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}
	fmt.Println(c.Query)
	fmt.Println("====================")

	// example access pod
	for _, pod := range c.Query.Pods {
		fmt.Println(pod.Title)
		for _, subpod := range pod.Subpods {
			fmt.Println("sub", subpod.Title)
		}
	}
	return
}
