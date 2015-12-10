package main
import "fmt"

type student struct{
	name string
	age int
}

func main(){
	vinit := student{name: "Vinit", age:23}
	nikita := student{name:"Nikita"}

	fmt.Println(vinit.name, vinit.age)
	fmt.Println(nikita.name)

	nikita.age=22
	fmt.Println(nikita.age)
}