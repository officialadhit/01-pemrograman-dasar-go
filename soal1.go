package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	// Menerima input nilai a, b, dan c
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)

	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	fmt.Print("Masukkan nilai c: ")
	fmt.Scan(&c)

	// Memeriksa apakah a^2 + b^2 == c^2
	if (a*a)+(b*b) == (c * c) {
		fmt.Println("pytagoras")
	} else {
		fmt.Println("not pytagoras")
	}
}
