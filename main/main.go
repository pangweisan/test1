package main

import (
	"fmt"
	"test1/pkg1"
)


func main(){

	fmt.Println("In main of module test1")
	mystr := pkg1.Hello()
	fmt.Println(mystr)
}

