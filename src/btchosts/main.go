package main

import (
	"fmt"
	"btchosts/hosts"
)

func main() {
	hosts.DoIt()
	fmt.Println("Press any key to exit...")
	fmt.Scanln()
}
