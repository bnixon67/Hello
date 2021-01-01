package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello Go World!")

	for n := 0; n < len(os.Args); n++ {
		fmt.Printf("Args[%d] = %s\n", n, os.Args[n])
	}
}
