package main

import (
	"fmt"
)

func main() {
	var harian, pekanan, akhir float64

	// Menerima input dengan teks petunjuk sesuai contoh
	fmt.Print("masukkan nilai evaluasi harian: ")
	fmt.Scan(&harian)

	fmt.Print("masukkan nilai evaluasi pekanan: ")
	fmt.Scan(&pekanan)

	fmt.Print("masukkan nilai evaluasi akhir: ")
	fmt.Scan(&akhir)

	// Validasi nilai input (harus di antara 0 dan 100)
	if harian < 0 || harian > 100 || pekanan < 0 || pekanan > 100 || akhir < 0 || akhir > 100 {
		fmt.Println("Invalid input")
		return
	}

	// Menghitung nilai akhir
	nilaiAkhir := (0.20 * harian) + (0.25 * pekanan) + (0.55 * akhir)

	// Menentukan kualifikasi hasil berdasarkan rentang nilai akhir
	switch {
	case nilaiAkhir >= 91:
		fmt.Println("Sangat Baik")
	case nilaiAkhir >= 80:
		fmt.Println("Baik")
	case nilaiAkhir >= 66:
		fmt.Println("Cukup")
	case nilaiAkhir >= 51:
		fmt.Println("Kurang")
	default:
		fmt.Println("Sangat Kurang")
	}
}
