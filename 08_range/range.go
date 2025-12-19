package main

import "fmt"

func main(){

//range helps us  to iterate over datastructures 

//using only for loop

nums :=[]int{3,4,5,2,4,5,3,5,6,4,6,4}

for i := 0; i < len(nums); i++ {
	fmt.Println(i+i*4);
}
for i,num :=range nums {
	fmt.Println(num,i)
}


//in maos
m :=map[string]string{"fname":"John Doe",
"mname":"Britto ","lname":"Mutthuswamy"};

for k,v :=range m {
	fmt.Println(k,v)
}

//unicode code point rune

for i,c := range "gollang" {
	fmt.Println(i,c)
}
}