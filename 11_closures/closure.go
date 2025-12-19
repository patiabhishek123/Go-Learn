package main

import "fmt"

//closures help us to preserve the value of a particular
// value which normally gets removed from the call stack

func main() {
 incrememt :=counter()
 fmt.Println(incrememt())
 fmt.Println(incrememt())
}
func counter() func() int {
	var count int =0;

	return func() int {
		count+=1
		return count
	}
}
