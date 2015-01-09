package main

import (
	"fmt"
	wolfram "github.com/KeizoBookman/go-wolfram-alpha"
)

func main() {
	id := ""
	c := wolfram.New(id)
	q,err := c.Get("oosaka")
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}
	fmt.Println(q)

	// example access pod
	for _, pod := range q.Pods {
		fmt.Println(pod.Title)
		for _, subpod := range pod.Subpods {
			fmt.Println("sub", subpod.Title)
		}
	}
	return
}
