package main

import (
	"fmt"
	"math"
)

func main() {
	// declare vars
	var a, b int
	// assign
	a = 5
	b = 10

	//add
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)

	//sub
	d := a - b
	fmt.Printf("%d - %d = %d\n", a, b, d)

	//div
	e := float32(a) / float32(b)
	fmt.Printf("%d / %d = %.2f\n", a, b, e)

	//mul
	f := a * b
	fmt.Printf("%d * %d=%d\n", a, b, f)

	demo2()
}

func demo() {
	a := 2.4
	b := 1.6

	c := math.Pow(a, 2)
	fmt.Printf("math.Pow(a,2) = %.2f\n", a, c)

	c = math.Sin(a)
	fmt.Printf("Sin(%.2f) = %.2f\n", a, c)

	c = math.Cos(b)
	fmt.Printf("Cos(%.2f) = %.2f\n", b, c)

	c = math.Sqrt(a * b)
	fmt.Printf("Sqrt(%.2f * %.2f) = %.2f", a, b, c)

}

func demo2() {
	var a = 4
	fmt.Printf("a = %d\n", a)

	a = a + 1
	fmt.Printf("a + 1 = %d\n", a)

	a++
	fmt.Printf("a++ = %d\n", a)

}
