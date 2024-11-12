package main

import (
	"fmt"
	"os"
)

func main() {
	var operasi int
	var operan1, operan2, hasil int
	fmt.Print("PROGRAM KALKULATOR SEDERHANA DI GOLANG\n1. OPERASI PERKALIAN\n2. OPERASI PEMBAGIAN\n3. OPERASI PENAMBAHAN\n4. OPERASI PENGURANGAN\n5. OPERASI MODULUS\n6. EXIT\nPilih salah satu angka di atas: ")
	fmt.Scan(&operasi)

	if operasi == 1 {
		fmt.Println("Perkalian")
		fmt.Print("Masukkan operan 1 : ")
		fmt.Scan(&operan1)

		fmt.Print("Masukkan operan 2 : ")
		fmt.Scan(&operan2)
		hasil = operan1 * operan2
		fmt.Println("Hasil : ", hasil)
	} else if operasi == 2 {
		fmt.Println("Pembagian")
		fmt.Print("Masukkan operan 1 : ")
		fmt.Scan(&operan1)

		fmt.Print("Masukkan operan 2 : ")
		fmt.Scan(&operan2)
		hasil = operan1 / operan2
		fmt.Println("Hasil : ", hasil)
	} else if operasi == 3 {
		fmt.Println("Penambahan")
		fmt.Print("Masukkan operan 1 : ")
		fmt.Scan(&operan1)

		fmt.Print("Masukkan operan 2 : ")
		fmt.Scan(&operan2)
		hasil = operan1 + operan2
		fmt.Println("Hasil : ", hasil)
	} else if operasi == 4 {
		fmt.Println("Pengurangan")
		fmt.Print("Masukkan operan 1 : ")
		fmt.Scan(&operan1)

		fmt.Print("Masukkan operan 2 : ")
		fmt.Scan(&operan2)
		hasil = operan1 - operan2
		fmt.Println("Hasil : ", hasil)
	} else if operasi == 5 {
		fmt.Println("Modulus")
		fmt.Print("Masukkan operan 1 : ")
		fmt.Scan(&operan1)

		fmt.Print("Masukkan operan 2 : ")
		fmt.Scan(&operan2)
		hasil = operan1 % operan2
		fmt.Println("Hasil : ", hasil)
	} else if operasi == 6 {
		fmt.Println("EXIT")
		os.Exit(0)
	} else {
		fmt.Println("Input tidak valid, silakan masukkan angka antara 1 hingga 6.")
	}
}