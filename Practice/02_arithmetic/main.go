package main

import "fmt"

func main() {
	a := 10
	b := 3

	fmt.Println("足し算")
	fmt.Println(a, "+", b, "=", a+b)

	fmt.Println("引き算")
	fmt.Println(a, "-", b, "=", a-b)

	fmt.Println("掛け算")
	fmt.Println(a, "*", b, "=", a*b)

	fmt.Println("引き算")
	fmt.Println(a, "/", b, "=", a/b)

	fmt.Println("あまり")
	fmt.Println(a, "%", b, "=", a%b)

}