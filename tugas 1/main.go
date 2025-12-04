package main

import "fmt"

func main() {
	// Deklarasi variabel
	var tugas, uts, uas, kehadiran float64
	var nilaiAkhir float64
	var grade string
	var status string

	// Input data
	fmt.Println("=== Program Penentu Nilai & Status ===")
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
	// Menggunakan if-else if untuk rentang nilai sesuai aturan
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

	// 3. Tentukan Status Kelulusan
	// Wajib menggunakan operator logika (&& atau ||) sesuai instruksi
	// Logika: Lulus jika Grade A/B/C DAN Kehadiran >= 75
	if (grade == "A" || grade == "B" || grade == "C") && kehadiran >= 75 {
		status = "Lulus"
	} else if (grade == "C" || grade == "D") && (kehadiran >= 60 && kehadiran < 75) {
		// Logika: Lulus Bersyarat jika Grade C/D DAN kehadiran antara 60-75
		status = "Lulus Bersyarat"
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

Kasus 1:
Input: Tugas=90, UTS=85, UAS=95, Kehadiran=80
Output: NA=90.5, Grade=A, Status=Lulus
Cabang IF terpakai:
- Grade: if nilaiAkhir >= 85 (Grade A)
- Status: if (grade A/B/C) && kehadiran >= 75 (Lulus)

Kasus 2:
Input: Tugas=60, UTS=60, UAS=60, Kehadiran=65
Output: NA=60, Grade=C, Status=Lulus Bersyarat
Cabang IF terpakai:
- Grade: else if nilaiAkhir >= 60 (Grade C)
- Status: else if (grade C/D) && (kehadiran >= 60...) (Lulus Bersyarat)

Kasus 3:
Input: Tugas=40, UTS=40, UAS=40, Kehadiran=90
Output: NA=40, Grade=E, Status=Tidak Lulus
Cabang IF terpakai:
- Grade: else (Grade E)
- Status: else (Tidak Lulus)

Pertanyaan: Apa yang terjadi kalau batas grade A diubah jadi 90?
Jawab: Maka mahasiswa dengan nilai 85-89 yang sebelumnya dapat A akan turun menjadi B.
*/