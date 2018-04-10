package main

import (
	"com/hosts"
	"fmt"
)

func main() {
	hosts.DoIt()
	fmt.Println("Press any key to exit...")
	fmt.Scanln()
}
