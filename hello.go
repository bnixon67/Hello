// Copyright 2024 Bill Nixon. All rights reserved.
// Use of this source code is governed by the license found in the LICENSE file.

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello")

	executable, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("executable %q\n", executable)

	for n, arg := range os.Args {
		fmt.Printf("args[%d] %q\n", n, arg)
	}
}
