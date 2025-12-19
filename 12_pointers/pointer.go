package main

import "fmt"
func changeNum(num int) {
	// in this nly the copy is received
	//hence no change in the original address.
	num = 5;
	fmt.Println("In changeNum",num)
}

func changeNumAdd(num *int) {
	*num=5
	fmt.Println("Value of num in changeAdd",num)
}
func main(){
	//simple pass by value where a copy of 
	// is passed
 num:=1
 changeNum(num)
 fmt.Println("After change Num",num)


   // using pointers
   changeNumAdd(&num);
   fmt.Println("Value of num after ChnageAdd",num)
}