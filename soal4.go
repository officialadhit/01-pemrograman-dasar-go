package main

import (
	"fmt"
)

func main() {
	var n int

	// Menerima input n
	fmt.Scan(&n)

	// Rumus jumlah tanaman: (n * n) / 2
	jumlahTanaman := (n * n) / 2

	// Menampilkan hasil
	fmt.Println(jumlahTanaman)
}
