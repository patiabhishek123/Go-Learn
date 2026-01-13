package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processNum(numchan chan int){
for num:= range numchan {
		fmt.Println("Processing channel -",num)
	time.Sleep(time.Second)}
}
//msin is processing itself in aseperate go routine
func main(){
	//channels are used to comminucate between goroutines
	numChan :=make(chan int)
	
	go processNum(numChan)
	for{
		numChan<-rand.Intn(100)
	}
	// numChan <-5


}