package main

import "fmt"

const a = 14

func main() {
	var b = 16
	var c = int64(20)

	d := a + b
	e := a + c
	fmt.Println("d:", d)
	fmt.Println("e:", e)

}
