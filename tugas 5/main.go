package main

import "fmt"

func main() {
	var tidur, air, screenTime int
	var skor int = 0
	var kategori string

	fmt.Println("=== Daily Habit Score ===")
	fmt.Print("Jam Tidur (jam): ")
	fmt.Scan(&tidur)
	fmt.Print("Air Minum (gelas): ")
	fmt.Scan(&air)
	fmt.Print("Screen Time (jam): ")
	fmt.Scan(&screenTime)

	// Hitung Skor Tidur
	// Tidak pakai else if berantai antar kategori karena skor dijumlah
	if tidur >= 7 && tidur <= 9 {
		skor += 40
	} else if tidur >= 5 && tidur <= 6 {
		skor += 20
	} else {
		skor += 5
	}

	// Hitung Skor Air Minum
	if air >= 8 {
		skor += 40
	} else if air >= 4 {
		skor += 20
	} else {
		skor += 5
	}

	// Hitung Skor Screen Time
	if screenTime <= 4 {
		skor += 20
	} else if screenTime <= 8 {
		skor += 10
	} else {
		skor += 0
	}

	// Tentukan Kategori Akhir
	if skor >= 80 {
		kategori = "Kebiasaan Sehat"
	} else if skor >= 50 {
		kategori = "Perlu Perbaikan"
	} else {
		kategori = "Bahaya (ubah pola hidup)"
	}

	fmt.Printf("\n--- Hasil Kesehatan ---\n")
	fmt.Printf("Total Skor : %d\n", skor)
	fmt.Printf("Kategori   : %s\n", kategori)
}

/*
================ KASUS UJI (TEST CASES) ================

Kasus 1:
Input: Tidur=8 (40), Air=8 (40), Screen=3 (20)
Output: Skor=100, Kategori=Kebiasaan Sehat
Cabang IF: Masuk ke kondisi poin tertinggi di semua kategori.

Kasus 2:
Input: Tidur=6 (20), Air=5 (20), Screen=6 (10)
Output: Skor=50, Kategori=Perlu Perbaikan
Cabang IF: Masuk ke kondisi tengah (else if) untuk semua input.

Kasus 3:
Input: Tidur=4 (5), Air=2 (5), Screen=10 (0)
Output: Skor=10, Kategori=Bahaya
Cabang IF: Masuk ke kondisi terendah (else) untuk semua input.

Pertanyaan: Apa yang terjadi jika poin tidur 7-9 jam dikurangi jadi 10?
Jawab: Skor maksimal akan turun drastis, sehingga lebih sulit mencapai kategori "Kebiasaan Sehat" (>=80).
*/