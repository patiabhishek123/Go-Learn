package main

import "fmt"

func main() {

	// fixed size,that is predictable
	// memory optimaixed
	// Constanr tme access
	var arr[4]int

	arr_with_initialization:= [3]int{34,23,31}
	// // fmt.Println(len(arr));

	// //2-D array
	arr3:=[2][2]int{{3,3},{4,5}}
	fmt.Println(arr)
	fmt.Println(arr_with_initialization)
	fmt.Println(arr3)


	//slices
	//dynamic arrays 
	//most userd construct

	//uninitialized slices is nill
	var nums[] int
	if nums==nil {
		fmt.Println("yes")
	} else {
		fmt.Println("No")
	}
	fmt.Println(len(nums))

	//make function is ised to add an empty array of some
	//  size in the existing slice

	var nums2= make([]int,4,6)
	fmt.Println(nums2)
	//capcity is the maximum number of 
	//element the alice can fit

	fmt.Println(cap(nums2));
	 

	//apenf function
	nums2=append(nums2,1)

}