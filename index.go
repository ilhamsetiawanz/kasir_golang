package main

import (
	"fmt"
	"kasir_golang/models"
	"kasir_golang/utils"
)

func main() {
	var belanjaan []models.Barang
	for {

		// Barang
		var nama string
		var harga float64
		var jumlah int

		// Memasukan Nama Barang
		fmt.Println("Masukan Nama Barang")
		fmt.Scanln(&nama)
		if nama == "Selesai" || nama == "Sudah" {
			break
		}
		// Memasukan Harga Barang
		fmt.Println("Masukan Harga Barang")
		fmt.Scanln(&harga)
		// Memasukan Jumlah Barang
		fmt.Println("Masukan Jumalah Barang")
		fmt.Scanln(&jumlah)

		item := models.Barang{
			Nama:   nama,
			Harga:  harga,
			Jumlah: jumlah,
		}

		belanjaan = append(belanjaan, item)
	}
	totalHarga := utils.CountJumlah(belanjaan)
	fmt.Printf("Total Harga Belanjanaan Anda Adalah Rp.%.2f", totalHarga)
}
