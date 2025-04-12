package main

import (
	"fmt"
	"goTmp/architectures"
	"goTmp/routers"
)

func main() {

	fmt.Println(architectures.MakeBase())
	fmt.Println(architectures.MakeLayered())
	fmt.Println(routers.MakeRouter("auth", 1))
	
}