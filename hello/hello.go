package main

import "fmt"

func main() {
	var str string
	var n, m int
	var mn float32

	str = "Hello 世界"
	n = 10
	m = 50
	mn = 2.45

	fmt.Println("value of str = ", str)
	fmt.Println("value of n= ", n)
	fmt.Println("value of m= ", m)
	fmt.Println("value of mn= ", mn)

	country := "China"
	val := 15
	fmt.Println("value of country= ", country)
	fmt.Println("value of val= ", val)

	var (
		name  string
		email string
		age   int
	)
	name = "John"
	email = "John@email.com"
	age = 23
	//
	/*formate some thing*/

	fmt.Println(name)
	fmt.Println(email)
	fmt.Println(age)

}
