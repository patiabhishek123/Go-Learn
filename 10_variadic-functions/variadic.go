package main

import "fmt"

//varidic functions are the functions
// which ca take any number of arguments


func sum(nums...int)int{
 //nums is receives as a slice of same datatype
	total:=0;
	for _,num:=range nums {
		total+=num;
	}
	return total;
}
// for using arguments of different datatypes we use interface()

func main() {
	result:=sum(5,3,4,5)
	result2:=[]int{3,5,2,6,7};
	//to pass slice of number we use smt similar to spread operator in js

  	fmt.Println(result);
	fmt.Println(sum(result2...))
}
 