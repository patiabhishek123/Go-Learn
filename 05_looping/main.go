package main

import "fmt"


func main(){

	//while loop
	i:=1;
	for i<=10 {
		fmt.Println(i)
		i=i+1;
	}

	//infinite loop
	// for {
	// 	println("a*")
	// }

	//Classic for loop
	for i:=0; i< 500;i++{
		
		if i%3==0 {
			continue
		}
		fmt.Println(i)
	}


	//1.22

	for i :=range 10{
		fmt.Println(i)
	}
}