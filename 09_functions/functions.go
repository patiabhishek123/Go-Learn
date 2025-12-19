package main

import "fmt"

func getSum(a int, b int)(int,int,bool){
	//we can return more ythan one values
	// in most cases we return the main return 
	// and the second value si mostly eror vlaue
	return a+b,a-b,a>b;
}

//we can also pass functions 

func main() {
	var a int;var b int;
	fmt.Scanln(&a);
	fmt.Scanln(&b)
	sum,diff,greater :=getSum(a,b)
	fmt.Println(getSum(a,b))
	fmt.Println(sum,diff,greater)
}
