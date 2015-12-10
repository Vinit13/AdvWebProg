package main
import "fmt"

func main(){
	var x uint = 50
	var y uint = 50
	fmt.Println("x is : ", x)
	fmt.Println("y is : ", y)
	x = x << 1
	y = y >> 1
	fmt.Println("x is now : ", x)
	fmt.Println("y is now : ", y)
}