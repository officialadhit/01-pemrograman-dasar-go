package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Menggunakan bufio.Scanner agar dapat membaca input yang mengandung spasi
	scanner := bufio.NewScanner(os.Stdin)

	// Menerima input kalimat
	fmt.Print("masukkan kalimat: ")
	scanner.Scan()
	kalimatInput := strings.TrimSpace(scanner.Text())

	// Menerima input kata kunci
	fmt.Print("masukkan kata kunci: ")
	scanner.Scan()
	kataKunciInput := strings.TrimSpace(scanner.Text())

	// Validasi jika salah satu atau kedua input kosong
	if kalimatInput == "" || kataKunciInput == "" {
		fmt.Println("Input invalid, harap isi kalimat dan kata kunci")
		return
	}

	// Mengubah seluruh huruf menjadi kecil (lowercase)
	kalimatLower := strings.ToLower(kalimatInput)
	kataKunciLower := strings.ToLower(kataKunciInput)

	// Memecah kalimat menjadi array/slice kata berdasarkan spasi
	daftarKata := strings.Split(kalimatLower, " ")

	// Mencari apakah kata kunci ada di dalam daftar kata
	ditemukan := false
	for _, kata := range daftarKata {
		if kata == kataKunciLower {
			ditemukan = true
			break
		}
	}

	// Menampilkan hasil pencarian
	if ditemukan {
		fmt.Println("kata ditemukan")
	} else {
		fmt.Println("kata tidak ditemukan")
	}
}
