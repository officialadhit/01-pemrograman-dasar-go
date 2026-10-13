package main

import (
	"fmt"
)

func main() {
	var n int

	// Menerima input bilangan
	fmt.Scan(&n)

	total := 0

	// Mencari faktor bilangan dari 1 hingga n dan menjumlahkannya
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			total += i
		}
	}

	// Menampilkan hasil penjumlahan faktor
	fmt.Println(total)
}
