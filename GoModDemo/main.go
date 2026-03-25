package main

import (
	"GoModDemo/app"
	"fmt"
)

func main() {
	s := hello()
	fmt.Println(s)
	c := app.NewComplex(1.0, 2.0)
	fmt.Println("test")
	fmt.Println(c.Real)
}
