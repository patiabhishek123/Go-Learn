package main

import (
	"fmt"
	"time"
)

func task(step int){
	fmt.Println("task number",step);
}


func main(){
	count  :=10;
	for i := 0; i < count; i++ {
		// go task(i)

		 go func(){
			fmt.Println(i);
		}()
	}

	time.Sleep(time.Second * 2)
}