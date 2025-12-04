package main

import "fmt"

func main() {
	var jenis, ukuran, jumlah int
	var hargaPerGelas, totalHarga int
	var kategori string

	fmt.Println("=== Coffee Order Checker ===")

	// 1. Validasi Jenis (Loop)
	// "for" tanpa kondisi = infinite loop (akan muter terus)
	for {
		fmt.Println("\nJenis: 1.Espresso, 2.Kopi Susu, 3.Chocolate")
		fmt.Print("Pilih Jenis (1-3): ")
		fmt.Scan(&jenis)

		// Cek validasi
		if jenis == 1 || jenis == 2 || jenis == 3 {
			break // KELUAR dari loop jika input benar
		}
		// Jika salah, loop akan ulang lagi dari atas
		fmt.Println("[!] Pilihan salah. Masukkan angka 1, 2, atau 3.")
	}

	// 2. Validasi Ukuran (Loop)
	for {
		fmt.Println("\nUkuran: 1.Small, 2.Medium, 3.Large")
		fmt.Print("Pilih Ukuran (1-3): ")
		fmt.Scan(&ukuran)

		if ukuran >= 1 && ukuran <= 3 {
			break
		}
		fmt.Println("[!] Ukuran salah. Pilih 1 (Small), 2 (Medium), atau 3 (Large).")
	}

	fmt.Print("\nJumlah Gelas: ")
	fmt.Scan(&jumlah)

	// Proses Hitung (Logic If-Else Utama)
	// Karena input sudah pasti valid (1-3) berkat loop diatas, logic bisa lebih bersih
	if jenis == 1 { // Espresso
		if ukuran == 1 {
			hargaPerGelas = 15000
		} else if ukuran == 2 {
			hargaPerGelas = 18000
		} else if ukuran == 3 {
			hargaPerGelas = 20000
		} else {
			fmt.Println("[!] Ukuran tidak valid.")
		}
	} else if jenis == 2 { // Kopi Susu
		if ukuran == 1 {
			hargaPerGelas = 18000
		} else if ukuran == 2 {
			hargaPerGelas = 22000
		} else if ukuran == 3 {
			hargaPerGelas = 25000
		} else {
			fmt.Println("[!] Ukuran tidak valid.")
		}
	} else { // Chocolate (Pasti jenis 3)
		if ukuran == 1 {
			hargaPerGelas = 17000
		} else if ukuran == 2 {
			hargaPerGelas = 20000
		} else if ukuran == 3 {
			hargaPerGelas = 23000
		} else {
			fmt.Println("[!] Ukuran tidak valid.")
		}
	}

	totalHarga = hargaPerGelas * jumlah

	// Tentukan Kategori
	if totalHarga > 40000 {
		kategori = "Pelanggan Sultan"
	} else if totalHarga >= 20000 {
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

Kasus 1 (Alur Normal - Hemat):
Input: Jenis=1 (Espresso), Ukuran=1 (Small), Jumlah=1
Output: Total=15000, Kategori=Pelanggan Hemat
Penjelasan: User langsung memasukkan data yang benar. Masuk ke if jenis==1 -> if ukuran==1.

Kasus 2 (Boundary - Batas Kategori):
Input: Jenis=3 (Chocolate), Ukuran=2 (Medium), Jumlah=2
Output: Total=40000, Kategori=Pelanggan Normal
Penjelasan: 
- Harga satuan: 20.000 (Choco Medium).
- Total: 20.000 * 2 = 40.000.
- Masuk kategori "Pelanggan Normal" (karena logika: 20.000 <= total <= 40.000).

Kasus 3 (Uji Validasi Loop - Input Salah Dulu):
Input (Urutan Eksekusi): 
1. Jenis  : 5 (Input Salah) -> Program menolak & minta ulang
2. Jenis  : 2 (Kopi Susu - Benar) -> Lanjut ke Ukuran
3. Ukuran : 4 (Input Salah) -> Program menolak & minta ulang
4. Ukuran : 3 (Large - Benar) -> Lanjut ke Jumlah
5. Jumlah : 2
Output: Total=50000, Kategori=Pelanggan Sultan
Penjelasan: 
- Program berhasil menahan user di dalam loop sampai input valid (1-3).
- Perhitungan akhir menggunakan input valid terakhir: Kopi Susu Large (25.000) * 2 = 50.000.
- Bagian "else { [!] Ukuran tidak valid }" di logika hitung tidak akan pernah tersentuh (Dead Code) karena input sudah dijaga oleh Loop.
*/