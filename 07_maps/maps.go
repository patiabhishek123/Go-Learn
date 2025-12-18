package main

import (
	"fmt"
	"maps"
)

//maps->hashtables
func main(){
	m:=make(map[string]string)
	m["abhishek"]="is a kind boy"
	m["adyasha"]="is abad girl"

	fmt.Println(m["abhishek"],m["adyasha"])

	//if the key value does not exist
	fmt.Println(m["babu"])
	// lengh of the map
	fmt.Println(len(m))
	// delete a key-value in amap
	delete(m,"adyasha")
	fmt.Println(len(m))
	//clear entire map
	clear(m)

	// map without make functuon

	m1:=map[string]int{"price":40,"phones":4}
	fmt.Println(m1)

	//check if an element is present is not or not
	_, ok:=m["price"]
	if ok {
		fmt.Println("all ok")

	}else{
		fmt.Println("something is wrong")
	}
	//equaltiy of maps
	m2:=map[string]int{"price":40,"color":001}
	m3:=map[string]int{"price":40,"color":001}
	m4:=map[string]int{"price":50,"color":01}
	fmt.Println(maps.Equal(m2,m3))
	fmt.Println(maps.Equal(m1,m4))

}	
