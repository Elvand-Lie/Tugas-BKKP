package main

import "fmt"

func main() {
	var jenis, ukuran, jumlah int
	var hargaPerGelas, totalHarga int
	var kategori string

	fmt.Println("=== Coffee Order Checker ===")
	fmt.Println("Jenis: 1.Espresso, 2.Kopi Susu, 3.Chocolate")
	fmt.Print("Pilih Jenis (1-3): ")
	fmt.Scan(&jenis)
	fmt.Println("Ukuran: 1.Small, 2.Medium, 3.Large")
	fmt.Print("Pilih Ukuran (1-3): ")
	fmt.Scan(&ukuran)
	fmt.Print("Jumlah Gelas: ")
	fmt.Scan(&jumlah)

	// 1. Hitung Harga Per Gelas
	// Menggunakan nested if untuk menentukan harga spesifik
	if jenis == 1 { // Espresso
		if ukuran == 1 {
			hargaPerGelas = 15000
		} else if ukuran == 2 {
			hargaPerGelas = 18000
		} else {
			hargaPerGelas = 20000
		}
	} else if jenis == 2 { // Kopi Susu
		if ukuran == 1 {
			hargaPerGelas = 18000
		} else if ukuran == 2 {
			hargaPerGelas = 22000
		} else {
			hargaPerGelas = 25000
		}
	} else if jenis == 3 { // Chocolate
		if ukuran == 1 {
			hargaPerGelas = 17000
		} else if ukuran == 2 {
			hargaPerGelas = 20000
		} else {
			hargaPerGelas = 23000
		}
	} else {
		hargaPerGelas = 0 // Salah input
	}

	// 2. Hitung Total
	totalHarga = hargaPerGelas * jumlah

	// 3. Tentukan Kategori Pelanggan
	// Menggunakan operator logika && untuk rentang
	if totalHarga > 40000 {
		kategori = "Pelanggan Sultan"
	} else if totalHarga >= 20000 && totalHarga <= 40000 {
		kategori = "Pelanggan Normal"
	} else {
		kategori = "Pelanggan Hemat"
	}

	fmt.Printf("\n--- Struk Pesanan ---\n")
	fmt.Printf("Total Harga : Rp %d\n", totalHarga)
	fmt.Printf("Kategori    : %s\n", kategori)
}

/*
================ KASUS UJI (TEST CASES) ================

Kasus 1:
Input: Jenis=1 (Espresso), Ukuran=1 (Small), Jumlah=1
Output: Total=15000, Kategori=Pelanggan Hemat
Cabang IF: if jenis==1 -> if ukuran==1, lalu else (total < 20rb)

Kasus 2:
Input: Jenis=2 (Kopi Susu), Ukuran=2 (Medium), Jumlah=2
Output: Harga=22000 * 2 = 44000, Kategori=Pelanggan Sultan
Cabang IF: else if jenis==2 -> else if ukuran==2, lalu if total > 40rb

Kasus 3:
Input: Jenis=3 (Choco), Ukuran=2 (Medium), Jumlah=1
Output: Total=20000, Kategori=Pelanggan Normal
Cabang IF: else if jenis==3 -> else if ukuran==2, lalu else if (20rb-40rb)

Pertanyaan: Apa yang terjadi jika harga Espresso Small naik jadi 21.000?
Jawab: Maka pemesan 1 gelas Espresso Small akan langsung masuk kategori "Pelanggan Normal" bukan lagi "Hemat".
*/