package main

import "fmt"

func main() {
	// Deklarasi variabel
	var tugas, uts, uas, kehadiran float64
	var nilaiAkhir float64
	var grade string
	var status string

	// Input data
	fmt.Println("=== Program Penentu Nilai & Status (Revisi) ===")
	fmt.Print("Masukkan Nilai Tugas (0-100): ")
	fmt.Scan(&tugas)
	fmt.Print("Masukkan Nilai UTS (0-100): ")
	fmt.Scan(&uts)
	fmt.Print("Masukkan Nilai UAS (0-100): ")
	fmt.Scan(&uas)
	fmt.Print("Masukkan Persentase Kehadiran (0-100): ")
	fmt.Scan(&kehadiran)

	// 1. Hitung Nilai Akhir (NA)
	nilaiAkhir = (0.3 * tugas) + (0.3 * uts) + (0.4 * uas)

	// 2. Tentukan Grade
	if nilaiAkhir >= 85 {
		grade = "A"
	} else if nilaiAkhir >= 70 {
		grade = "B"
	} else if nilaiAkhir >= 60 {
		grade = "C"
	} else if nilaiAkhir >= 50 {
		grade = "D"
	} else {
		grade = "E"
	}

	// 3. Tentukan Status Kelulusan (SOLUSI LOGIKA DIPERBAIKI)
	// Aturan 1: Lulus Murni (Nilai A-C dan Kehadiran aman)
	if (grade == "A" || grade == "B" || grade == "C") && kehadiran >= 75 {
		status = "Lulus"

	// Aturan 2: Lulus Bersyarat (REVISI)
	// Kita tambahkan grade "A" dan "B" disini.
	// Jadi jika grade A/B/C/D tapi kehadiran hanya 60-74, masuknya Lulus Bersyarat.
	} else if (grade == "A" || grade == "B" || grade == "C" || grade == "D") && (kehadiran >= 60 && kehadiran < 75) {
		status = "Lulus Bersyarat"

	// Aturan 3: Sisanya gagal
	} else {
		status = "Tidak Lulus"
	}

	// Output
	fmt.Printf("\n--- Hasil ---\n")
	fmt.Printf("Nilai Akhir : %.2f\n", nilaiAkhir)
	fmt.Printf("Grade       : %s\n", grade)
	fmt.Printf("Status      : %s\n", status)
}

/*
================ KASUS UJI (TEST CASES) ================

Kasus 1 (Lulus Murni):
Input: Tugas=90, UTS=85, UAS=95, Kehadiran=80
Output: NA=90.5, Grade=A, Status=Lulus

Kasus 2 (Kasus yang tadinya Error/Tidak Lulus):
Input: Tugas=78, UTS=55, UAS=90, Kehadiran=61
Output: NA=75.9, Grade=B, Status=Lulus Bersyarat
Penjelasan: 
- Sebelumnya outputnya "Tidak Lulus" karena Grade B tidak dicek di blok else-if.
- Sekarang sudah dicek (grade A/B/C/D) && (kehadiran 60-75), jadi statusnya benar.

Kasus 3 (Tidak Lulus):
Input: Tugas=40, UTS=40, UAS=40, Kehadiran=90
Output: NA=40, Grade=E, Status=Tidak Lulus
*/