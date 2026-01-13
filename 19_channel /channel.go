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

func emailSender(emailChan<- chan string,done chan <-bool ){
	defer func ()  {
			done<- true
		}()
	for email:= range emailChan{
		fmt.Println("sending email to : ",email)
		time.Sleep(time.Second)
		
	}
}
//main is processing itself in aseperate go routine

func main(){
	//channels are used to comminucate between goroutines
	// numChan :=make(chan int)
	// go processNum(numChan)
	// for{
	// 	numChan<-rand.Intn(100)
	// }
	// // numChan <-5

	emailChan:=make(chan string)
	done:=make(chan bool)
	go emailSender(emailChan,done)
	for i := 0; i < 100; i++ {
		emailChan<-fmt.Sprint("%d%dabhishek@gmail.com",i,rand.Int31n(50));
	}
	fmt.Println("message is sending ... ")




}