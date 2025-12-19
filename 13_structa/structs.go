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
func main (){
	myorder :=order{
		id:      "32",
		amount:  50.00,
		status:  "received",
		createdAt: time.Now(),
		
	}
	fmt.Println("Order struct",myorder)
}
