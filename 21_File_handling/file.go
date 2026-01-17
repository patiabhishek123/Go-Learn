package main

import (
	"fmt"
	"os"
)

// import (
//     "fmt"
//     "io/ioutil"
//     "log"
//     "os"
// )

// func main() {
//     // Write to a file
//     content := "Hello, World!"
//     err := ioutil.WriteFile("output.txt", []byte(content), 0644)
//     if err != nil {
//         log.Fatal(err)
//     }
//     fmt.Println("File written successfully")

//     // Read from a file
//     data, err := ioutil.ReadFile("output.txt")
//     if err != nil {
//         log.Fatal(err)
//     }
//     fmt.Println("File content:", string(data))

//     // Append to a file
//     file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0644)
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer file.Close()

//     _, err = file.WriteString("\nAppended text")
//     if err != nil {
//         log.Fatal(err)
//     }
//     fmt.Println("Content appended successfully")
// }


func main (){
	f,err:=os.Open("output.txt")
	if err != nil{
		panic(err)
	}
	fmt.Println("file info ",f.Name())
}