package main

import (
	"fmt"
	
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


	//EMAIL SENDER
	// emailChan:=make(chan string)
	// done:=make(chan bool)
	// go emailSender(emailChan,done)
	// for i := 0; i < 10; i++ {
	// 	x:=rand.Int31n(50)
	// 	emailChan<-fmt.Sprintf("%dabhishek@gmail.com",x);
	// }
	// fmt.Println("message is sending ... ")
	// close(emailChan)

	//HANDLING MULTIPLE CHANNELS
	chan1:=make(chan int)
	chan2:=make(chan string)

	go func(){
		chan1<-69
	}()
	go func(){
		chan2<-"pong"
	}()

	for i := 0; i < 2; i++ {
		select{
		case chan1Val:=<-chan1:
			fmt.Printf("%d......",chan1Val)
		
		case chan2Val:=<-chan2:
			fmt.Printf("%s......",chan2Val)
		}
	}



}