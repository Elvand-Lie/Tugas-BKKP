package main

import "fmt"

func main() {
	var jamMasuk, jamKeluar, tipeMember int
	var durasi int
	var tarifDasar, totalBayar float64
	var diskon float64

	fmt.Println("=== Parkir Smart ===")
	fmt.Print("Jam Masuk (0-23): ")
	fmt.Scan(&jamMasuk)
	fmt.Print("Jam Keluar (0-23): ")
	fmt.Scan(&jamKeluar)
	fmt.Println("Tipe Member (0=Non, 1=Silver, 2=Gold): ")
	fmt.Scan(&tipeMember)

	if jamKeluar < jamMasuk {
		fmt.Printf("Jam keluar (%d) tidak boleh kurang dari jam masuk (%d).\n", jamKeluar, jamMasuk)
		return
	}

	// 1. Hitung Durasi
	durasi = jamKeluar - jamMasuk
	// Catatan: Kode "if durasi < 0" yang lama sudah tidak perlu 
	// karena sudah dicek oleh validasi di atas.

	// 2. Tentukan Tarif Dasar
	if durasi <= 2 {
		tarifDasar = 5000
	} else if durasi <= 5 {
		tarifDasar = 10000
	} else {
		tarifDasar = 15000
	}

	// 3. Terapkan Diskon
	if tipeMember == 2 { // Gold
		diskon = 0.50 * tarifDasar
	} else if tipeMember == 1 { // Silver
		diskon = 0.20 * tarifDasar
	} else {
		diskon = 0
	}

	totalBayar = tarifDasar - diskon

	fmt.Printf("\n--- Tagihan Parkir ---\n")
	fmt.Printf("Durasi Parkir : %d Jam\n", durasi)
	fmt.Printf("Tarif Dasar   : Rp %.0f\n", tarifDasar)
	fmt.Printf("Diskon        : Rp %.0f\n", diskon)
	fmt.Printf("Total Bayar   : Rp %.0f\n", totalBayar)
}

/*
================ KASUS UJI (TEST CASES) ================

Kasus 1 (Normal):
Input: Masuk=8, Keluar=10, Member=0 (Non)
Output: Durasi=2, Tarif=5000, Diskon=0, Total=5000

Kasus 2 (Diskon Gold):
Input: Masuk=10, Keluar=14, Member=2 (Gold)
Output: Durasi=4, Tarif=10000, Diskon=5000, Total=5000

Kasus 3 (Diskon Silver):
Input: Masuk=7, Keluar=15, Member=1 (Silver)
Output: Durasi=8, Tarif=15000, Diskon=3000, Total=12000

Kasus 4 (Validasi Error - BARU):
Input: Masuk=10, Keluar=8 (Logika Terbalik)
Output: [ERROR] Jam keluar (8) tidak boleh kurang dari jam masuk (10).
Penjelasan: Program langsung berhenti (return) dan tidak mencetak tagihan.
*/