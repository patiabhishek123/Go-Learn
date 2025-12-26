package main

import "fmt"
func main(){
	intarr :=[] int{2,3,4,5};
	stringarr :=[] string{"Abhishke","Loves " ,"Adyasha"}
printSlice(intarr);
printSlice(stringarr);


///stack
stk:=[]int{1,2,3,4,5};
myStack :=stack[int]{
	elements: stk,
}
fmt.Println(myStack)

}

// how to use genenrics in struct and use them

type stack[T any] struct{
	elements []int
}

func printSlice[T any](items []T){
	for _,i := range items{
	fmt.Println(i);

	}
	
	}
	
