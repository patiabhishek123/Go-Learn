package main

import (
	"fmt"
	"time"
)

//order struct

type order struct {
	id 			string
	amount      float32
	status	    string
	createdAt   time.Time 
}

//receiver for structs similar to implement
//method in rust
func (o *order) chnageStatus(status string) {
	//sutomatic derefrencing is done by structs itself
	o.status=status;
}

func main (){
	myorder :=order{
		id:      "32",
		amount:  50.00,
		status:  "received",
		createdAt: time.Now(),
		
	}
	fmt.Println("Order struct",myorder)
	fmt.Println("Status before changStatus ,",myorder.status)
	myorder.chnageStatus("waiting");
	fmt.Println("Status after change",myorder.status)
}
