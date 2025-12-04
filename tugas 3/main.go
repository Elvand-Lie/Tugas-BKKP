package main

import "fmt"

func main() {
	var budget int
	var kebutuhan int
	var paket string

	fmt.Println("=== Rekomendasi Paket Internet ===")
	fmt.Print("Masukkan Budget Anda (Rp): ")
	fmt.Scan(&budget)
	fmt.Println("Pilih Kebutuhan Utama:")
	fmt.Println("1. Chat & Sosmed")
	fmt.Println("2. Streaming Video")
	fmt.Println("3. Gaming Online")
	fmt.Print("Pilihan (1-3): ")
	fmt.Scan(&kebutuhan)

	// Logika Percabangan Bersarang (Nested IF)
	if kebutuhan == 1 {
		// Kebutuhan Chat & Sosmed
		if budget < 50000 {
			paket = "Paket Sosmed Hemat (2GB)"
		} else if budget <= 100000 {
			paket = "Paket Sosmed Standar (10GB + Unli Chat)"
		} else {
			paket = "Paket Sosmed Sultan (Unlimited)"
		}
	} else if kebutuhan == 2 {
		// Kebutuhan Streaming
		if budget < 75000 {
			paket = "Paket Nonton Hemat (5GB)"
		} else if budget <= 150000 {
			paket = "Paket Nonton Puas (20GB + Youtube)"
		} else {
			paket = "Paket Bioskop 4K (100GB)"
		}
	} else if kebutuhan == 3 {
		// Kebutuhan Gaming
		// Menggunakan operator logika OR untuk variasi
		if budget < 50000 || budget == 0 {
			paket = "Paket Game Mobile (3GB + Ping Stabil)"
		} else {
			paket = "Paket Pro Gamer (50GB + Priority)"
		}
	} else {
		paket = "Kebutuhan tidak valid"
	}

	fmt.Printf("\n--- Rekomendasi ---\n")
	fmt.Printf("Kebutuhan : %d\n", kebutuhan)
	fmt.Printf("Budget    : Rp %d\n", budget)
	fmt.Printf("Paket     : %s\n", paket)
}

/*
================ KASUS UJI (TEST CASES) ================

Kasus 1:
Input: Budget=30000, Kebutuhan=1 (Chat)
Output: Paket Sosmed Hemat (2GB)
Cabang IF: if kebutuhan == 1 -> if budget < 50000

Kasus 2:
Input: Budget=100000, Kebutuhan=2 (Streaming)
Output: Paket Nonton Puas (20GB + Youtube)
Cabang IF: else if kebutuhan == 2 -> else if budget <= 150000

Kasus 3:
Input: Budget=200000, Kebutuhan=3 (Gaming)
Output: Paket Pro Gamer (50GB + Priority)
Cabang IF: else if kebutuhan == 3 -> else (budget besar)

Pertanyaan: Apa yang terjadi jika kategori Gaming dihapus?
Jawab: Input '3' akan masuk ke kondisi 'else' terakhir (Kebutuhan tidak valid) atau program perlu diubah logikanya.
*/