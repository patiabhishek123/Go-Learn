package main

import "fmt"
//we cannot use shorthand notation 
// outside the function it gives error 
// eg .  name:="ram"
// const x="dog" this is absoutely legal 

func main() {
	var a int = 23
	fmt.Println(a)
	var name string = "golang"
	var name2 = "golang infernce without using var"
	fmt.Println(name)
    fmt.Println(name2)

	//shorthand syntax for assigning and declaring variabe simutaneously
	animal:="tiger"
	fmt.Println(animal)
}
