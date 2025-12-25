package main

import "fmt"

type OrderStatus int;
const(
	Received OrderStatus=iota
	Delivered
	Sent
	OnTheWay
)
// func OrderStatusReceived(status int){
// 	fmt.Println("Order Status , ",status);
// }

func ChangeOrderStatus(status OrderStatus){
	fmt.Println("changeing order status to ", status)
}
func main(){
	ChangeOrderStatus(OnTheWay)

	
}