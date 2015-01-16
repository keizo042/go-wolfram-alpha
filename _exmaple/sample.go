package main
/*
this sample is simple, minimam, shortest.
when you try this sample, you get Wolfram Developer id &
declear i variable.
like this.
id := YOUR_WOLFRAM_ID.

That way, this sample take correct work.

*/

import (
	"fmt"
	wolfram "github.com/KeizoBookman/go-wolfram-alpha"
)

func main() {
	id := ""
	c := wolfram.New(id)
	q,err := c.Get("oosaka", wolfram.A_PLAIN)
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
