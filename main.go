package main

import (
	"fmt"
	"os"
)



func main(){
	setup()
	args := os.Args
	if args[1] == "all" {
		getter()
		reader()
		//indexer()
	} else {
		fmt.Println("Usage: ./vortigaunt all")
	}
	}





