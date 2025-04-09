package main

import (
	"fmt"
	
	configfile "goTmp/config_file"
)

func main() {
	fmt.Println(configfile.MakeConfig(1, "gin_db", "12345678", "postgres"))
	// fmt.Println(architectures.MakeBase())
	// fmt.Println(architectures.MakeLayered())
	// fmt.Println(architectures.MakeSqlc(1, "internal/db", "goTmp"))
}