package main
import "fmt"

func main () {
	myString := "Hello World"
	for i:=0; i < len(myString); i++ {
		fmt.Println (myString[i], " - ", string(myString[i]))
	}
}