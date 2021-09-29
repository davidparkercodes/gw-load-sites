package main

import (
	"flag"
	"fmt"
)

func main() {

	aaa := flag.String("example", "defaultValue", " Help text.")
	flag.Parse()

	fmt.Println(*aaa)
}
