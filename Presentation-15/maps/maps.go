package main
import "fmt"

func main(){
	myMap := map[string]string{
		"Vinit":"Nair",
		"Ajay":"Bir",
		"SG":"K",
	}

	myMap["HSHK"] = "BlahBlah"
	myMap["KHSH"] = "halBhalB"
	delete(myMap, "KHSH")
	for key, val := range myMap{
		fmt.Println(key, " - ", val)
	}

	fmt.Println(len(myMap))
	if val, ok := myMap["Ajay"]; ok{
		fmt.Println(val)
	}
}