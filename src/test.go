package main

import (
	"fmt"
	"pkg1"
	"pkg2"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(pkg1.ReturnStr())
	fmt.Println(pkg1.Pack1Int)
	fmt.Println(pkg1.Pack1Float)
	fmt.Println(pkg2.ReturnStr())
}
