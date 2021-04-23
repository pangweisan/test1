package pkg2

import (
	"fmt"
	"test1/pkg1"
)


func Hello2(){

	fmt.Println("In Hello2 of module test1")
	mystr := pkg1.Hello()
	fmt.Println(mystr)
}

