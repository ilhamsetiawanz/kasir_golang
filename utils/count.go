package utils

import "kasir_golang/models"

func CountJumlah(items []models.Barang) float64 {
	var Hasil float64
	for _, item := range items {
		Hasil += item.Harga * float64(item.Jumlah)
	}
	return Hasil
}
